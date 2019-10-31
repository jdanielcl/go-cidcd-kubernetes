pipeline {
    agent any
    environment {
        PATH = "/usr/local/go/bin:$PATH"
        DOCKER_IMAGE_NAME = "jdanielcl/go-cicd-kubernetes"
        DOCKER_CONTAINER_NAME = "go-cidcd-kubernetes"
    }
    stages {
        stage('Build') {
            steps {
                echo 'Compiling Program'
                def rtDocker = Artifactory.server 'GoogleJFrog'
           
                // Attach custom properties to the published artifacts:
                rtDocker.addProperty("project-name", DOCKER_IMAGE_NAME).addProperty("status", "stable")
            
                // Push a docker image to Artifactory (here we're pushing hello-world:latest). The push method also expects
                // Artifactory repository name (<target-artifactory-repository>).
                // Please make sure that <artifactoryDockerRegistry> is configured to reference the <target-artifactory-repository> Artifactory repository. In case it references a different repository, your build will fail with "Could not find manifest.json in Artifactory..." following the push.
                def buildInfo = rtDocker.push '${DOCKER_IMAGE_NAME}:latest', DOCKER_CONTAINER_NAME
            
                // Publish the build-info to Artifactory:
                server.publishBuildInfo buildInfo
            }
        }        
        stage('JFrog artifactory') {
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
        stage('Run docker container'){
            steps{
                sh 'docker stop ${DOCKER_CONTAINER_NAME}'
                sh 'docker rm ${DOCKER_CONTAINER_NAME}'
                sh 'docker run --name ${DOCKER_CONTAINER_NAME} -d -p 85:8181 ${DOCKER_IMAGE_NAME}'
            }
        }
    }
}