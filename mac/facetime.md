# faceTime

添加自动接受

    $ defaults write com.apple.FaceTime AutoAcceptInvites -bool YES

添加自动接受的上账号

    $ defaults write com.apple.FaceTime AutoAcceptInvitesFrom -array-add erasinoo@gmail.com

自动接听电话

    $ defaults write com.apple.FaceTime AutoAcceptInvitesFrom -array-add +13588452919
