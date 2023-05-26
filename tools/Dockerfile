FROM animcogn/face_recognition:cpu
# 使用Python 3.11.3作为基础镜像

# 设置工作目录
WORKDIR /facerecognition

# 复制当前目录下的所有文件到容器中的/app目录下
COPY facerecognition /facerecognition


# 暴露端口
EXPOSE 43221

# 运行Python程序
CMD ["python3", "server.py"]