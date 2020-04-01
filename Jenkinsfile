pipeline {
    agent {label '!master'}
    stages { 
        stage('Build') {
            steps {
                sh "docker rm -f new-server || true"
                sh 'go build -o server ./src'
                sh "docker pull ubuntu:18.04"
            }
        }
        stage('Run') {
            steps {
                sh "docker run -d --rm --name new-server -w /app -v $WORKSPACE:/app ubuntu:18.04 /bin/bash -c '/app/server'"
            }
        }
    }
}