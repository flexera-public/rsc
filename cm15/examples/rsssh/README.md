# rsssh

This is a simple project that makes use of RightScale API using rsc and builds a simple bash aliases file with SSH
links to various machines in various environments.

The input configuration looks like this:

```json
{
  "SshUser": "rightscale",
  "SshOptions": "-i ~/.id_rsa_rightscale -o StrictHostKeyChecking=no",
  "OutputFile": "/Users/arangamani/.rightscale_aliases",
  "Environments": {
    "dev": {
      "Account": 12345,
      "Servers": {
        "lb": "Loadbalancer",
        "dbmaster": "Database Manager Master",
        "dbslave": "Database Manager Slave"
      },
      "ServerArrays": {
        "app": "Application Server"
      }
    },
    "prod": {
      "Account": 12346,
      "Servers": {
        "lb": "Loadbalancer",
        "dbmaster": "Database Manager Master",
        "dbslave": "Database Manager Slave"
      },
      "ServerArrays": {
        "app": "Application Server"
      }
    }
  }
}
```

The output is in bash aliases file format:

```
alias dev_lb='ssh -i ~/.id_rsa_rightscale -o StrictHostKeyChecking=no rightscale@<IP Address>'
alias dev_dbmaster='ssh -i ~/.id_rsa_rightscale -o StrictHostKeyChecking=no rightscale@<IP Address>'
alias dev_dbmaster='ssh -i ~/.id_rsa_rightscale -o StrictHostKeyChecking=no rightscale@<IP Address>'
alias dev_app#1='ssh -i ~/.id_rsa_rightscale -o StrictHostKeyChecking=no rightscale@<IP Address>'
alias prod_lb='ssh -i ~/.id_rsa_rightscale -o StrictHostKeyChecking=no rightscale@<IP Address>'
alias prod_dbmaster='ssh -i ~/.id_rsa_rightscale -o StrictHostKeyChecking=no rightscale@<IP Address>'
alias prod_dbmaster='ssh -i ~/.id_rsa_rightscale -o StrictHostKeyChecking=no rightscale@<IP Address>'
alias prod_app#1='ssh -i ~/.id_rsa_rightscale -o StrictHostKeyChecking=no rightscale@<IP Address>'
alias prod_app#2='ssh -i ~/.id_rsa_rightscale -o StrictHostKeyChecking=no rightscale@<IP Address>'
alias prod_app#3='ssh -i ~/.id_rsa_rightscale -o StrictHostKeyChecking=no rightscale@<IP Address>'
```

Add this to your bash profile:
```bash
if [ -f ~/.rightscale_aliases ]; then
    . ~/.rightscale_aliases
fi
```

Run it using
```
rsssh -e $RS_USER_EMAIL -p $RS_USER_PASSWORD
```

By default, the configuration file is read from `~/.rsssh`, but it can be overridden using the `-c` flag.

