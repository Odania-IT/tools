# AWS Bootstrap

This can setup basic roles for terraform usage in AWS.

The configuration is in a yaml file. You can apply it by calling it like

    ./generate.rb CONFIG_FILE

You need a config like:

    ---
    variables:
      main_account_id: MAIN-ACCOUNT-ID
      main_profile: MAIN-PROFILE
      region: eu-central-1
      intialize_terraform: true
      state_bucket_name: STATE-BUCKET-NAME
      state_dynamo_table: STATE-DYNAMODB-NAME
    
    accounts:
      develop:
        account_id: 123
        profile_name: your-organisagion-profile-develop
      test:
        account_id: 345
        profile_name: your-organisagion-profile-test
      live:
        account_id: 456
        profile_name: your-organisagion-profile-live
      infrastructure:
        account-id: 678
        profile_name: your-organisagion-profile-infrastructure
    
    groups:
      developer:
        managed-policies:
          - arn:aws:iam::aws:policy/ReadOnlyAccess
      readonly:
        managed-policies:
          - arn:aws:iam::aws:policy/ReadOnlyAccess
      admin:
        managed-policies:
          - arn:aws:iam::aws:policy/ReadOnlyAccess
      org-admin:
        managed-policies:
          - arn:aws:iam::aws:policy/AdministratorAccess
    
    managed_policies:
      admin:
        statements:
          - effect: Allow
            resources:
              - '*'
            actions:
              - '*'
      billing-readonly:
        statements:
          - effect: Allow
            resources:
              - '*'
            actions:
              - aws-portal:View*
              - budgets:View*
      developer:
        statements:
          - effect: Allow
            resources:
              - '*'
            actions:
              - acm:*
              - apigateway:*
              # dynamoDB
              - application-autoscaling:DeleteScalingPolicy
              - application-autoscaling:DeregisterScalableTarget
              - application-autoscaling:DescribeScalableTargets
              - application-autoscaling:DescribeScalingActivities
              - application-autoscaling:DescribeScalingPolicies
              - application-autoscaling:PutScalingPolicy
              - application-autoscaling:RegisterScalableTarget
              - appsync:*
              - autoscaling:*
              - cognito-idp:*
              - cognito-identity:*
              - cloudformation:*
              - cloudfront:*
              - cloudwatch:*
              - codebuild:*
              - codecommit:*
              - codepipeline:*
              - dynamodb:*
              - ec2:*
              - ecr:*
              - elasticfilesystem:*
              - elasticloadbalancing:*
              - events:*
              - kms:*
              - lambda:*
              - logs:*
              - s3:*
              - sns:*
              - iam:*
              - cloudtrail:*
              - trustedadvisor:*
              - route53:*
              - route53domains:*
              - ssm:*
              - firehose:*
              - rds:*
    
    roles:
      admin:
        policies:
          - admin
          - developer
          - billing-readonly
      developer:
        policies:
          - developer
          - billing-readonly
      readonly:
        policies:
          - billing-readonly
    
    users:
      max.mustermann:
        groups:
          - admin
          - developer
          - org-admin
          - readonly

This config will create the appropriate roles in every account. The users and groups will be created in the main account.
Each user can assume the other accounts with the roles he owns.
