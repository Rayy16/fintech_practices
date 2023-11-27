# config ssh-server & check server status
sudo apt-get install openssh-server
ps -e | grep ssh

# install python3.8 with apt
sudo apt install python3.8

# link to /usr/bin/python
sudo ln -s /usr/bin/python3.8 /usr/bin/python3.8

# install pip with get-pip.py
curl https://bootstrap.pypa.io/get-pip.py -o get-pip.py 
sudo python get-pip.py

# may be we need link to /usr/bin/pip
which pip3
sudo ln -s /usr/bin/pip3 /usr/bin/pip

# upgrade pip module
python -m pip install pip --upgrade

# install virtualenv module
pip install virtualenv

# build virtual environment
mkdir envs && cd ./envs
virtualenv sadtalker && virtualenv mockingbird

# init environment with requeirement.txt
cd ..
source ./envs/sadtalker/bin/activate
pip install ffmpeg
pip install -r ./SadTalker/requirement.txt
deactivate

source ./envs/mockingbird/bin/activate
pip install ffmpeg
sudo apt-get install python3.8-dev
pip install -r ./MockingBird/requeirement.txt
deactivate