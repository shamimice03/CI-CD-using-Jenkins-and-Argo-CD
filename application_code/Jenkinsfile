pipeline {
    agent any
    environment{
        DOCKERHUB_CREDS = credentials('dockerhub')
    }
    stages {
        stage('Clone Repo') {
            steps {
                checkout scm
                sh 'ls *'
            }
        }
        stage('Build Image') {
            steps {
                sh 'docker build -t shamimice03/go-greetings-app:$BUILD_NUMBER ./application_code '
            }
        }
        stage('Docker Login') {
            steps {
                sh 'echo $DOCKERHUB_CREDS_PSW | docker login -u $DOCKERHUB_CREDS_USR --password-stdin'
            }
        }
        stage('Docker Push') {
            steps {
                sh 'docker push shamimice03/go-greetings-app:$BUILD_NUMBER'
            }
        }
        stage('Trigger Linked Task') {
            steps {
                build job: 'webapp-linked-one', parameters: [string(name: 'DOCKERTAG', value: env.BUILD_NUMBER )]
            }
        }
    }
    post {
        always {
            sh 'docker logout'
        }
    }
}
