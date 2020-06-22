pipeline {
  agent  any
  stages {
    stage('verify branch'){
      steps {
        echo "BRANCH: {$GIT_BRANCH}"
      }
    }
    stage('Build and Push'){
      steps {
        withCredentials([string(credentialsId: 'DB_CONN_STR', variable: 'DB_CONN_STR')]) {
          script {
            docker.withRegistry('','DockerHub'){
              def image = docker.build('chaturvedisulabh/go-web-services-101:${env.BUILD_ID}', '--build-arg -DB_CONN_STR=$DB_CONN_STR')
              image.push()
            }
          }
        }
      }
    }
    stage('Run'){
      steps {
        withCredentials([string(credentialsId: 'DB_CONN_STR', variable: 'DB_CONN_STR')]) {
          docker.image('chaturvedisulabh/go-web-services-101:${env.BUILD_ID}').withRun('--name go-web-services-101 -p 5000:5000 -PORT=5000 -DB_CONN_STR=$DB_CONN_STR')
        }
      }
    }
  }
}