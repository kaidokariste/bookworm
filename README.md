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

## UNIX > Enimkasutatavad käsud
`$touch myFile`    - uus fail  
`$pwd`             - kausta teekond  
`$mkdir myFolder` - uue kataloogi loomine  
`$rmdir myFolder` - tühja kataloogi kustutamine  
`$rm -rf myFolder` - kataloogi kustutamine kui ta sisaldab faile  
`$ps -A` - kõik masinas jooksvad protsessid  
`$top` - aktiivsed protsessid  
`$ps -u myUser` - kasutajaspetsiifiliselt aktiivsed protsessid  
`$sudo cp /home/kaido/database.php /var/www/html/config` - faili/kausta kopeerimine  
`$ls -lah` - faili täiendatud nimekiri  
`$cat myFile` - faili sisu vaatamine  
`$sudo tar -jcvf myArchive.tbz2 myFolder` - arhiivi tekitamine kaustast  
`$tar -jxvf myArchive.tbz2` - arhiivi lahtipakkimine  
`$du -h` - disk usage  
`$du -h -d 1` - disk usage üks kategooria täpsemalt  
`$df -h` - üldine kettakasutus  
`$mv myOldFileName myNewFileName` - faili ümbernimetamine  
`$ln -s git/myProject symApp` - loob sümlingi täpsustatud kaustast aktiivsesse kausta, kaustanimega symApp  
`$realpath --relative-to=/folder/from /folder/to` - suhtelise kausta tee saamine (folder/from tavaliselt projekti root)

## UNIX > SSH ühendus remote serverisse
 - *myComputer* - arvuti, millega soovime repote serverisse/arvutisse ühendada
 - *remoteHost* - arvuti millesse ühendume

*myComputer -> remoteHost*

 1. Genereeri myComputeris võtmepaar _private_key_ ja _public_key_ ```$ssh-keygen -t rsa 2048```  
 2. remoteHostis ```$vim .ssh/authorized_keys``` ja lisa sinna juurde genereeritud _public_key_
 3. Puttyga ühendades, anna ette _myUsername@remoteHost_ ja lisa juurde _private_key_
 
# Git
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

# Docker
## Enimkasutatavad käsud
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

### Weblinks

1. [Connecting to PostgreSQL database with Go](https://www.calhoun.io/using-postgresql-with-golang/)
