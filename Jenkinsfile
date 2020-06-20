pipeline {
  agent  any
  stages {
    stage('verify branch'){
      steps {
        echo "BRANCH: {$GIT_BRANCH}"
      }
    }
    stage('Docker Build and Push'){
      steps {
        sh '''
        docker build -t go-web-services-101 .
        docker push chaturvedisulabh/go-web-services-101
        '''
      }
    }
    stage('Docker Run'){
      steps {
        sh '''
        docker run --name go-web-services-101 -p 5000:5000 chaturvedisulabh/go-web-services-101 -port=5000 -DB_CONN_STR=
        '''
      }
    }
  }
}