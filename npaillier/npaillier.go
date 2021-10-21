package npaillier

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
)

//var zero = new(big.Int).SetInt64(0)
var one = new(big.Int).SetInt64(1)

type PublicKey struct {
	N  *big.Int // n
	N2 *big.Int // n^2
	g  *big.Int // 生产源，对于使用私有私钥解密的时候需要
	h  *big.Int // 独立公钥
}

func (pk *PublicKey) Add(c1 *CipherPub, c2 *CipherPub) *CipherPub {
	c := new(CipherPub)
	c.T1 = new(big.Int).Mul(c1.T1, c2.T1)
	c.T1.Mod(c.T1, pk.N2)
	c.T2 = new(big.Int).Mul(c1.T2, c2.T2)
	c.T2.Mod(c.T2, pk.N2)
	return c
}

// 私有私钥，可以用来解密自己的Pk加密内容
type PrivateKey struct {
	Pk    *PublicKey
	sigma *big.Int // 私有密钥
}

// 密文，如果是使用私有私钥就需要T2才能解密，如果使用lambda则不需要T2
type CipherPub struct {
	T1 *big.Int
	T2 *big.Int
}

func (sk PrivateKey) GetPk() *PublicKey {
	return sk.Pk
}

func (pk *PublicKey) Encrypt(msg *big.Int) *CipherPub {
	r := getRandom(pk.N2)
	c := new(CipherPub)
	t := new(big.Int).Mul(msg, pk.N)
	t.Add(t, one)
	t.Mul(t, new(big.Int).Exp(pk.h, r, pk.N2))
	t.Mod(t, pk.N2)
	c.T1 = t
	c.T2 = new(big.Int).Exp(pk.g, r, pk.N2)
	return c
}

func (sk *PrivateKey) Decrypt(c *CipherPub) *big.Int {
	pk := sk.Pk
	u := new(big.Int).Exp(c.T2, sk.sigma, pk.N2)
	u.ModInverse(u, pk.N2)
	u = L(u.Mul(c.T1, u), pk.N)
	u.Mod(u, pk.N)
	return u
}

type NPPrivatekey struct {
	pk  *PublicKey    // 联合公钥
	sks []*PrivateKey // 公私钥对，这些公钥加密的也能用lambda解密

	x *big.Int // 联合公钥的私钥

	lambda  *big.Int // lambda也是私钥，这是全局私钥，可以解密任何公钥
	vlambda *big.Int // lambda逆元
	lambda1 *big.Int // lambda拆分私钥1
	lambda2 *big.Int // lambda拆分私钥2
}

/*
	获取联合公钥
*/
func (sk NPPrivatekey) GetPk() *PublicKey {
	return sk.pk
}

/*
	获取私钥对
*/
func (sk NPPrivatekey) GetSks() []*PrivateKey {
	return sk.sks
}

// 生成NPPrivatekey
func GenkeyPair(bitlen int, beta int) *NPPrivatekey {
	if bitlen < 1024 {
		log.Fatalf("The `bitlen` parameter should not be smaller then 1024")
		return nil
	}
	p := getPrime(bitlen / 2)
	q := getPrime(bitlen / 2)
	a := getPrime(bitlen / 2)
	x := getPrime(bitlen / 2)
	n := new(big.Int).Mul(p, q)
	n2 := new(big.Int).Mul(n, n)
	g := new(big.Int).Mod(new(big.Int).Neg(new(big.Int).Exp(a, new(big.Int).Mul(n, big.NewInt(2)), n2)), n2)
	h := new(big.Int).Exp(g, x, n2)
	pk := &PublicKey{
		N:  n,
		N2: n2,
		g:  g,
		h:  h,
	}

	lambda := phi(q, p)
	vlambda := new(big.Int).ModInverse(lambda, n2)
	S := new(big.Int).Mul(lambda, vlambda)
	S.Mod(S, new(big.Int).Mul(lambda, n2))
	lambda1 := getPrime(bitlen)
	lambda2 := S.Sub(S, lambda1)

	sk := &NPPrivatekey{
		pk:      pk,
		x:       x,
		lambda:  lambda,
		vlambda: vlambda,
		lambda1: lambda1,
		lambda2: lambda2,
		sks:     make([]*PrivateKey, beta),
	}

	for i := 0; i < beta; i++ {
		xi := getPrime(bitlen - 12)
		sk.sks[i] = &PrivateKey{
			Pk:    &PublicKey{n, n2, g, new(big.Int).Exp(g, xi, n2)},
			sigma: xi,
		}
	}
	return sk
}

// 联合公钥的私钥解密
func (sk *NPPrivatekey) DecryptByx(c *CipherPub) *big.Int {
	pk := sk.pk
	u := new(big.Int).Exp(c.T2, sk.x, pk.N2)
	u.ModInverse(u, pk.N2)
	u = L(u.Mul(c.T1, u), pk.N)
	u.Mod(u, pk.N)
	return u
}

// 联合私钥直接解密
func (sk *NPPrivatekey) Decrypt(T1 *big.Int) *big.Int {
	pk := sk.pk
	u := new(big.Int).Exp(T1, sk.lambda, pk.N2)
	u = L(u, pk.N)
	u.Mul(u, sk.vlambda)
	u.Mod(u, pk.N)
	return u
}

func (sk *NPPrivatekey) DecryptLambda1(T1 *big.Int) *big.Int {
	pk := sk.pk
	return new(big.Int).Exp(T1, sk.lambda1, pk.N2)
}

func (sk *NPPrivatekey) DecryptLambda2(T1 *big.Int) *big.Int {
	pk := sk.pk
	return new(big.Int).Exp(T1, sk.lambda2, pk.N2)
}

func (sk *NPPrivatekey) DecryptBy2(T1 *big.Int) *big.Int {
	pk := sk.pk
	u1 := sk.DecryptLambda1(T1)
	u2 := sk.DecryptLambda2(T1)
	u1.Mul(u1, u2)
	u1.Mod(u1, pk.N2)
	u1 = L(u1, pk.N)
	return u1.Mod(u1, pk.N)
}

// L (x,n) = (x-1)/n is the largest integer quocient `q` to satisfy (x-1) >= q*n
func L(x, n *big.Int) *big.Int {
	return new(big.Int).Div(new(big.Int).Sub(x, one), n)
}

// generates a random number, testing if it is a probable prime
func getPrime(bits int) *big.Int {
	p, err := rand.Prime(rand.Reader, bits)
	if err != nil {
		panic("Error while reading crypto/rand")
	}

	return p
}

// getRandom generates a random Int `r` such that `r < n` and `gcd(r,n) = 1`
func getRandom(n *big.Int) *big.Int {
	gcd := new(big.Int)
	r := new(big.Int)
	err := fmt.Errorf("")

	for gcd.Cmp(one) != 0 {
		r, err = rand.Int(rand.Reader, n)
		if err != nil {
			panic("Error while reading crypto/rand")
		}

		gcd = new(big.Int).GCD(nil, nil, r, n)
	}
	return r
}

// Computes Euler's totient function `φ(p,q) = (p-1)*(q-1)`
func phi(x, y *big.Int) *big.Int {
	p1 := new(big.Int).Sub(x, one)
	q1 := new(big.Int).Sub(y, one)
	return new(big.Int).Mul(p1, q1)
}
