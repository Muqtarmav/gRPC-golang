pipeline {
    agent any
     tools {
        go 'go1.14'
    }
    
     environment {
         GO114MODULE = 'on'
         CGO_ENABLED = 0 
         GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
     }
    stages {        
        stage('Pre Test') {
            steps {
                echo 'Installing dependencies'
                sh 'go version'
                sh 'go get -u golang.org/x/lint/golint'
            }
        }
    
        stage('docker build'){
            steps {
                script{
                sh 'docker build -t muqtarmav/go-roava:3.0 .' 
            }
        }
        
    }
        
         stage('Push image to Hub'){
            steps{
                script{
                   withCredentials([string(credentialsId: 'dockerhub', variable: 'dockerhub')]) {
                   sh 'docker login -u muqtarmav -p ${dockerhub}'
}
                   sh 'docker push muqtarmav/go-roava:2.0'
                }
            }
        }
        
    }   
} 
    
