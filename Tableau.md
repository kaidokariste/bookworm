![](img/tableaulogo.png )

* [Tableau commands](#Tableau-releases)
* [Purge server](#Purge-server)
* [Installimine 2025.3.4 näitel](#Installimine-2025.3.4-näitel)

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

# Installimine 2025.3.4 näitel
