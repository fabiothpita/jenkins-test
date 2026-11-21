pipeline {
	agent any

	stages {
		stage('Checkout') {
			steps {
				echo 'Checking out...'
				checkout scm
			}
		}
		stage('Test') {
			steps {
				echo 'Testing...'
				sh '''
					docker build \
					--target test \
					-t helpdesk-api:test .
				'''
			}
		}
		stage('Build') {
			steps {
				echo 'Building...'
				sh ''' 
					docker build \
					-t helpdesk-api:latest .
				'''
			}
		}
		stage('Deploy') {
			steps {
				echo 'Deploying...'
				sh ''' 
					docker rm -f helpdesk-api || true

					docker run -d \
					--name helpdesk-api \
					-p 127.0.0.1:8081:8081 \
					helpdesk-api:latest
				
				'''
			}
		}
		stage('smoke test') {
			step {
			
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