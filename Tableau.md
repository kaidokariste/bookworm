![](img/tableaulogo.png )

* [Tableau releases](#Tableau-releases)
* [Purge server](#Purge-server)
* [TSM põhikäsud](#TSM-põhikäsud)

# Tableau releases
[Server releases](https://www.tableau.com/support/releases/server)

# Purge server
```bash
# Kontrolli, kas Tableau pakett on installitud
dpkg -l | grep tableau

#Vastavalt paketi versioonile
sudo apt-get purge tableau-server-20253.26.0109.0333
sudo rm -rf /var/opt/tableau
sudo rm -rf /etc/opt/tableau
sudo rm -rf /opt/tableau
sudo apt autoremove -y
sudo apt clean
```

# TSM põhikäsud
```bash
curl -x <cach-server-if needed> http://downloads.tableau.com/tssoftware/tableau-server-2025-3-4_amd64.deb -o tableau-server-2025-3-4_amd64.deb
sudo apt-get install ./tableau-server-2025-3-4_amd64.deb - Ubuntus tableau install.
tsm licenses list - litsentside vaatamine
tsm licenses activate -k <SINU_VÕTI> - litsentsi aktiveerimine  
tsm settings import -f /var/opt/tableau/identity_store.json - AD konfiguratsioonijsoni importimine  
```

