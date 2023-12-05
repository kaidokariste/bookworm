# UNIX
## UNIX> Uue kasutaja loomine 
`$cat /etc/passwd`- kasutajate vaatamine  
`$sudo useradd -m -s /bin/bash myusername` - Kasutaja loomine  
`$sudo passwd myusername` - Parooli määramine  
`$sudo gpasswd -a myusername sudo` - Kasutaja lisamine sudo gruppi

## UNIX> Vim
Esc koolon ja ..  
`set number` - Näitab faili reanumbreid  
`set nonumber` - Eemaldab failist reanumbrid  
`set bg=dark` - Kaotab ära tumesinise teksti, muudab heledamaks  

## UNIX > Enimkasutatavad käsud
`$touch myFile`    - uus fail  
`$pwd`             - kausta teekond  
`$mkdir myFolder` - uue kataloogi loomine  
`$rmdir myFolder` - tühja kataloogi kustutamine  
`$rm -rf myFolder` - kataloogi kustutamine kui ta sisaldab faile  
`$ps -A` - kõik masinas jooksvad protsessid  
`$ps -aux | grep mongo`  - mongo protsessi otsimine kõikide protsesside hulgast  
`$top -i` - aktiivsed protsessid ja CPU kasutus  
`$ps -u myUser` - kasutajaspetsiifiliselt aktiivsed protsessid  
`$sudo cp /home/kaido/database.php /var/www/html/config` - faili/kausta kopeerimine  
`$ls -lah` - faili täiendatud nimekiri  
`$cat myFile` - faili sisu vaatamine  
`$sudo tar -jcvf myArchive.tbz2 myFolder` - arhiivi tekitamine kaustast  
`$tar -jxvf myArchive.tbz2` - arhiivi lahtipakkimine  
`$du -sh /folder/path/` - Kausta suurus GB (h)uman readable vormis.     
`$du -h --max-depth=1 html` - kausta html alamkaustade suurused   
`$df -h` - üldine kettakasutus  
`$mv myOldFileName myNewFileName` - faili ümbernimetamine  
`$ln -s git/myProject symApp` - loob uue sümlingi täpsustatud kaustast aktiivsesse kausta, kaustanimega symApp  
`$ln -sf git/myProject symApp` - loob või uuendab sümlingi täpsustatud kaustast aktiivsesse kausta, kaustanimega symApp  
`$realpath --relative-to=/folder/from /folder/to` - suhtelise kausta tee saamine (folder/from tavaliselt projekti root)  
`$cat /proc/meminfo` või `$cat /proc/meminfo` - Info masina mälu kohta . Viimases info MB  
`$find . -name "file_name.sh"` - Search some specific file  
`$find /var/www/html/user/files/ -type f -mtime +30 -ls` - vanemad kui 30 päeva failid koos lisainfoga
`$find /home/airflow_user/airflow/logs/scheduler -type f -mtime +30 -delete` - kustuta vanemad kui 30 päeva logi failid.  
`$find . -type d -empty -delete` - kustuta kausta sees olevad tühjad kaustad  
`$chown -R kaido:kaido /run/some.file` - grant file owner and owning group to kaido  

## UNIX > SSH ühendus remote serverisse
 - *myComputer* - arvuti, millega soovime repote serverisse/arvutisse ühendada
 - *remoteHost* - arvuti millesse ühendume

*myComputer -> remoteHost*

 1. Genereeri myComputeris võtmepaar _private_key_ ja _public_key_ ```$ssh-keygen -t rsa 2048```  
 2. remoteHostis ```$vim .ssh/authorized_keys``` ja lisa sinna juurde genereeritud _public_key_
 3. Puttyga ühendades, anna ette _myUsername@remoteHost_ ja lisa juurde _private_key_
 
# Git
## Intalling GIT
Ubuntu
```bash
$ sudo apt-get install git
git --version
```
CentOS
```
yum install git
git --version
```
Generat SSH key pair to connect repository (as local user)  
```
$ ssh-keygen -t rsa
```   
Copy your public key to add it into Github.
```
$ cat ~/.ssh/id_rsa.pub
```
Configure your global identity
```
$ git config --global user.name "firstname.surname"
$ git config --global user.email "firstname.surname@mydomain.xx"
```
Clone your repo
```
$ cd /your/desired/folder
$ git clone ssh://git@stash.com:7999/myrepo/myrepo.git
```
Check out desired branch
```
git checkout --track remotes/origin/master
```


## Git > Enimkasutataud käsud

`$git remote -v ` -näitab "fetch" ja "pull" remote aadresse  
`$git checkout --track remotes/myBranch.git` - uuele remote branchile ümberlülitumine koos muudatuste jälgimisega  
`$git status` - hetkel kasutatav branch  
`$git branch -a` - kõik olemasolevad branchid  
`$git branch -d myBranch` - kohaliku branchi kustutamine  
`$git branch -d $(git branch)` - kõikide kohalike branchide kustutamine  
`$git branch -r` - kõik remotebranchid  
`$git pull ` - muudatuste allalaadimine  
`$git clone https://github.com/myRepo` - repository kloonimine  
`$git add . ` - failide lisamine committimiseks  
`$git remote prune origin --dry-run` - kuvab remote branchide nimekirja, mida on võimalik tühjendada  
`$git remote prune origin` - tühjendab näidatavate remote branchide nimekirja  
`$find -name *.properties -exec git rm --cached {} \;`- Remove files from index but dont remove file itself . 

# Docker
`$docker images `- näitab teeki kohalikest Docker imagetest  
`$docker ps ` - aktiivsete konteinerite loend  
`$docker ps -a ` - kõik konteinerid  
`$docker rm <container-id>` - konteineri eemaldamine id kaudu  
`$docker rm $(docker ps -aq)` - kõikide konteinerite kustutamine  
`$docker rmi <image-id:version>` - dockerimage eemaldamine image-id põhjal  
`$docker rmi $(docker images -q) ` - image kustutamine hulga kaupa  
`$docker inspect <container-id>` - konteineri info vaatamine  
`$docker run -i -t <image-id> /bin/bash` - konteineri jooksutamine olemasolevast konteinerist  
`$docker start <conteiner-id> ` - seisva konteineri käivitamine  
`$docker stop <container-id>` - petab etteantud aktiivse konteineri  
`$docker exec -i -t <conteiner-id> /bin/bash` - jooksvasse konteinerisse sisenemine  
`$docker run -d <image-id>` - konteineri alustamine ja jooksutamine tagaplaanil  
`$docker run -d -P <image-id:tag>` - konteineri jooksutamine tagaplaanil ja lisaks kõikide konteineri vabade portide ühendamine host pordiga  
`$docker push <repository:tag>` - dockerhubi reposse üleslaadimine  
`$docker cp container:src_path dest:path ` - failide kopeerimine konteinerist hosti  
`$docker cp SRC_PATH CONTAINER:DEST_PATH ` - failide kopeerimine hostist konteinerisse  
`$docker build --no-cache -t <container-name> .` - ilma cachita konteineri buildimine dockerfailist  
`$docker compose rm -svf` - docker-compose.yaml olemasolul keskkonna puhastamine enne jooksutamist  
`$docker compose up -d` - docker-compose.yaml olemasolul teenuste konteinerite jooksutamine taustal  
`$docker logs <container_id>` - konteineri logid, töötab ka mitteaktiivse konteineri puhul  
`$sudo systemctl start docker` - starting Docker  
`$docker run -p 127.0.0.1:3000:5000 docker-image-tag`- bind docker container port 5000 to localhost port 3000  
`$docker run -v volume-name:/path/in/docker/image container-tag` - named volumes (choose your volume_name and docker creates this itself outside container).  
`$docker run -dp 5000:5000 -w /app -v "$(pwd):/app" flask-smorest-api` - Oma lokaalse repo mountimine dockeri konteinerisse. (UNIX süsteemid)  
`$docker volume inspect volume-name` - vaata, kus _volume_ asub host masinas.  

### Weblinks

1. [Connecting to PostgreSQL database with Go](https://www.calhoun.io/using-postgresql-with-golang/)
