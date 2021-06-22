# 基于java:8
FROM openjdk:11
# 复制当前目录下的order-0.0.1-SNAPSHOT.jar文件（这是springboot项目打包后的压缩包）到容器内
COPY app.jar /app.jar

CMD ["--server.port=8081"]
# 暴露端口8010
EXPOSE 8081
# 启动项目
ENTRYPOINT ["java","-jar","/app.jar"]