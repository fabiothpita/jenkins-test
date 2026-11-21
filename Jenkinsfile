pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Test') {
            steps {
                sh '''
                    docker build \
                      --target test \
                      -t helpdesk-api:test .
                '''
            }
        }

        stage('Build') {
            steps {
                sh '''
                    docker build \
                      -t helpdesk-api:latest .
                '''
            }
        }

        stage('Deploy') {
            steps {
                sh '''
                    docker rm -f helpdesk-api || true

                    docker run -d \
                      --name helpdesk-api \
                      -p 127.0.0.1:8081:8081 \
                      helpdesk-api:latest
                '''
            }
        }

        stage('Smoke Test') {
            steps {
                sh '''
                    docker exec helpdesk-api \
                      wget -qO- http://localhost:8081/health
                '''
            }
        }
    }

    post {
        success {
            echo 'Deploy realizado com sucesso!'
        }

        failure {
            echo 'Pipeline falhou!'
        }
    }
}