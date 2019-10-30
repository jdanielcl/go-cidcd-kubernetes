pipeline {
    agent any
    environment {
        DOCKER_IMAGE_NAME = "jdanielcl/go-cicd-kubernetes"
    }
    stages {
        stage('Build') {
            steps {
                echo 'Compiling Program'
                sh 'go test -v'
            }
        }
        stage('Build Docker Image') {
            steps {
                echo 'Compiling Program'
            }
        }
    }
}