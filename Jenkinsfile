pipeline {
    agent any

	parameters {
		choice(
			name: 'DEPLOY_ENV',
			choices: ['staging', 'production'],
			description: 'Ambiente de deploy'
		)

		booleanParam(
			name: 'RUN_TESTS',
			defaultValue: true,
			description: 'Executar testes unitários?'
		)
	}

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Test') {
			when {
				expression {
					params.RUN_TESTS == true
				}
			}
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

       
		stage('Deploy staging') {
			when{
				expression {
					params.DEPLOY_ENV == 'staging'
					params.RUN_TESTS == true
				}
			}

			steps {

				sh '''
				
					docker rm -f helpesk-staging || true
					
					docker run -d \
						--name helpdesk-staging \
						-p 127.0.0.1:8081:8081 \
						helpdesk-api:latest
				
				'''
			}
		}

		stage ('Deploy production') {

			when {
				expression {
					params.DEPLOY_ENV == 'production'
					params.RUN_TESTS == true
				}
			}

			steps {
				sh '''
					docker rm -f helpdesk-production || true

					docker run -d \
					--name helpdesk-production \
					-p 127.0.0.1:8082:8081 \
					helpdesk-api:latest
				
				 '''
			}
		}

        stage('Smoke Test Staging') {
			when {
				expression {
					params.DEPLOY_ENV == 'staging'
					params.RUN_TESTS == true
				}
			}
            steps {
                sh '''
                    docker exec helpdesk-staging \
                      wget -qO- http://localhost:8081/health
                '''
            }
        }


		stage('Smoke Test Production') {
			when {
				expression {
					params.DEPLOY_ENV == 'production'
					params.RUN_TESTS == true
				}
			}
            steps {
                sh '''
                    docker exec helpdesk-production \
                      wget -qO- http://localhost:8082/health
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