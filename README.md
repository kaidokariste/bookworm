# UNIX
## UNIX> Uue kasutaja loomine 
`$sudo useradd -m -s /bin/bash myusername` - Kasustaja loomine  
`$sudo passwd myusername` - Parooli määramine  
`$sudo gpasswd -a myusername sudo` - Kasutaja lisamine sudo gruppi

## UNIX > Enimkasutatavad käsud
`$touch myFile`    - uus fail  
`$pwd`             - kausta teekond  
`$mkdir myFolder` - uue kataloogi loomine  
```$rmdir myFolder``` - tühja kataloogi kustutamine  
```$rm -rf myFolder``` - kataloogi kustutamine kui ta sisaldab faile  
```$ps -A``` - kõik masinas jooksvad protsessid  
```$top``` - aktiivsed protsessid  
```$ps -u myUser``` - kasutajaspetsiifiliselt aktiivsed protsessid  
`$sudo cp /home/kaido/database.php /var/www/html/config` - faili/kausta kopeerimine  
```$ls -lah``` - faili täiendatud nimekiri  
```$cat myFile``` - faili sisu vaatamine  
```$sudo tar -jcvf myArchive.tbz2 myFolder``` - arhiivi tekitamine kaustast  
```$tar -jxvf myArchive.tbz2``` - arhiivi lahtipakkimine  
```$du -h``` - disk usage  
```$du -h -d 1``` - disk usage üks kategooria täpsemalt  
```$df -h``` - üldine kettakasutus  
```$mv myOldFileName myNewFileName``` - faili ümbernimetamine  
```$ln -s git/myProject symApp``` - loob sümlingi täpsustatud kaustast aktiivsesse kausta, kaustanimega symApp

## UNIX > SSH ühendus remote serverisse
 - *myComputer* - arvuti, millega soovime repote serverisse/arvutisse ühendada
 - *remoteHost* - arvuti millesse ühendume

*myComputer -> remoteHost*

 1. Genereeri myComputeris võtmepaar _private_key_ ja _public_key_ ```$ssh-keygen -t rsa 2048```  
 2. remoteHostis ```$vim .ssh/authorized_keys``` ja lisa sinna juurde genereeritud _public_key_
 3. Puttyga ühendades, anna ette _myUsername@remoteHost_ ja lisa juurde _private_key_

# Golang
## GoLang Types

**Numeric types**
```
uint8       the set of all unsigned  8-bit integers (0 to 255)
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

int8        the set of all signed  8-bit integers (-128 to 127)
int16       the set of all signed 16-bit integers (-32768 to 32767)
int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

float32     the set of all IEEE-754 32-bit floating-point numbers
float64     the set of all IEEE-754 64-bit floating-point numbers

complex64   the set of all complex numbers with float32 real and imaginary parts
complex128  the set of all complex numbers with float64 real and imaginary parts

byte        alias for uint8
rune        alias for int32
```

### Weblinks

1. [Connecting to PostgreSQL database with Go](https://www.calhoun.io/using-postgresql-with-golang/)
