# Prepare swamp setup

Swamp helps to switch between different AWS accounts. The helper creates bash aliases that make it easy to configure.

## Create a config

Here is an example config that sets up aliases for two profiles with different accounts. The account you login (profile)
should require mfa to secure the sub accounts. The sub accounts first login to the main account and afterwords switch
to the other account. The AWS_PROFILE enviroment variable is set to the "target_profile".

If you want to request some API-Keys or anything else from the aws you can do it with the extra_script section.

    ---
    profile1:
      mfa_device: 'arn:aws:iam::ACCOUNT_ID:mfa/MFA_ID'
      profile: 'base-profile-1'
      region: 'eu-central-1'
      target_role: 'user/developer'
      extra_script: |
        #!/usr/bin/env bash
        echo Setup anything special you need
      roles:
        - account: 'TARGET_ACCOUNT_ID_1'
          target_profile: 'profile1-account1-dev-developer'
        - account: 'TARGET_ACCOUNT_ID_2'
          target_profile: 'profile1-account2-prod-developer'
        - account: 'TARGET_ACCOUNT_ID_3'
          target_profile: 'profile1-account3-dev-other-role'
          target_role: 'user/OTHER_TARGET_ROLE'
    profile2:
      mfa_device: 'arn:aws:iam::ACCOUNT_ID:mfa/MFA_ID'
      profile: 'base-profile-2'
      region: 'eu-central-1'
      target_role: 'developer'
      roles:
        - account: 'TARGET_ACCOUNT_ID_1'
          target_profile: 'profile2-account1-dev-developer'

The extra scripts file is written to

	.swamp_extra_scripts.d

relative to your home folder. A file called PROFILE_scripts.sh is created there.

## Setup

To setup the aliases simply call:

    ./prepare-swamp.rb

or

    ./prepare-swamp.rb MY-CONFIG-FILE.yml

You need to confirm that the file will be written to your home dir under :

	.bash_aliases.d/swamp-aliases.sh

You need to add something like this to your .bashrc in your home folder:

	source ~/.bash_aliases.d/swamp-aliases.sh

## Swamp

The swamp tool needs to be installed. You can find releases here:

[https://github.com/felixb/swamp](https://github.com/felixb/swamp)
