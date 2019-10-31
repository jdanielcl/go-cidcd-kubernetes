pipeline {
    agent any
    environment {
        PATH = "/usr/local/go/bin:$PATH"
        DOCKER_IMAGE_NAME = "jdanielcl/go-cicd-kubernetes"
        DOCKER_CONTAINER_NAME = "go-cidcd-kubernetes"
    }
    stages {
        stage('Build Test') {
            steps {
                echo 'Artifactory JFrog'
                sh 'go test -v'
                sh 'echo $USER'
            }
        }
        stage('Build Docker Image') {
            steps {
                script {
                    app = docker.build(DOCKER_IMAGE_NAME)
                    app.withRun("-d -p 8181:8181") { c ->
                        sh 'curl localhost:8181'
                    }    
                }
            }
        }
        stage('Build Artifactory') {
            steps {
                sh 'echo $USER_JFROG'
                SH 'go mod init ${DOCKER_IMAGE_NAME}'
                sh 'curl -X PUT -$USER_JFROG:$PASSWORD_JFROG -T go.mod "http://104.198.64.55/artifactory/go-cdcd-kubernetes/"'
            }
        }    
        stage('Run docker container'){
            steps{
                sh 'docker stop ${DOCKER_CONTAINER_NAME}'
                sh 'docker rm ${DOCKER_CONTAINER_NAME}'
                sh 'docker run --name ${DOCKER_CONTAINER_NAME} -d -p 85:8181 ${DOCKER_IMAGE_NAME}'
            }
        }
    }
}