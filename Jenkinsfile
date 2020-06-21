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
        withCredentials([string(credentialsId: 'DB_CONN_STR', variable: 'DB_CONN_STR')]) {
          sh '''
            docker run -d --name go-web-services-101 -p 5000:5000 chaturvedisulabh/go-web-services-101:latest -PORT=5000 -DB_CONN_STR=$DB_CONN_STR
          '''
        }
      }
    }
  }
}