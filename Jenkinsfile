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
        sh '''
        docker run --name go-web-services-101 -p 5000:5000 chaturvedisulabh/go-web-services-101:latest -port=5000 -DB_CONN_STR=
        '''
      }
    }
  }
}