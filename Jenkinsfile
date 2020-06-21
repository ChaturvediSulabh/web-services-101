pipeline {
  agent  any
  stages {
    stage('verify branch'){
      steps {
        echo "BRANCH: {$GIT_BRANCH}"
      }
    }
    stage('Docker Build'){
      steps {
        sh '''
        docker build -t chaturvedisulabh/go-web-services-101:latest .
        '''
      }
    }
    stage('Docker Run'){
      steps {
        node {
        withCredentials([string(credentialsId: 'DB_CONN_STR', variable: 'POSTGRES DB CONNECTION STRING')])
        sh '''
          docker run --name go-web-services-101 -p 5000:5000 chaturvedisulabh/go-web-services-101:latest -PORT=5000 -DB_CONN_STR=$DB_CONN_STR
        '''
        }
      }
    }
  }
}