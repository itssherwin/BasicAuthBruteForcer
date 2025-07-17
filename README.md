# BasicAuthBruteForcer


Http Auth BruteForcer Tool Written in GoLang Supports both Basic and Digest authentication types

simple and fast http auth bruteforce

## ✨ Features
* **Ultra-fast**: Harnesses Go’s concurrency to brute force credentials with 50 threads by default (customizable).
* **Supports Both Auth Types**: Brute force HTTP **Basic** and HTTP **Digest** authentication.
* **Flexible Input**: Test single credentials or entire wordlists for users and passwords.
* **Timeout Control**: Tweak HTTP connection timeouts for slow or unreliable networks.
Colorful, Clear Output: Easy-to-read status, success, and failure messages.
* **Open Source**: MIT licensed and ready for your improvements.

## 🛠 Install
```bash
go install github.com/itssherwin/BasicAuthBruteForcer@latest
```


## 🚦 Usage

```bash
$ BasicAuthBruteForcer
        -user string.
                 One User to test.
        -passwd string.
                 One Password to test.
        -user-file FILE.
                 List of users.
        -passwd-file FILE.
                 List of passwords.
        -timeout int.
                 HTTP Connection timeout. Default: 10s
        -auth-type string.
                 Authentication type: basic or digest. Default: basic
        -threads int.
                 Number of concurrent threads (goroutines). Default: 50
```


### Examples

Brute force Digest Auth with custom timeout and 100 threads:
:
```bash
$ BasicAuthBruteForcer -url http://target -user-file users.txt -passwd-file passwords.txt -auth-type digest -timeout 30 -threads 100
```

Password spray:
```bash
$ BasicAuthBruteForcer -url https://target/protected -user-file users.txt -passwd 'Def@ultPassw0rd'
```

📦 Default wordlists are included in the wordlists/ directory:
* wordlists/users.txt
* wordlists/passwords.txt

Feel free to use your own!


## 📝 License
This project is licensed under the MIT.
Contributions, issues, and feature requests are welcome! 🚀

## Support 💎
Thanks for your support!
* Bitcoin
```
bc1qq6vrlnytva67mj956nydfyvuzwl4t6wy2naahc
```
* Ethereum
```
0xa88238491Df0219b0F924Fc6c6e1Bc8B3BB50E60
```
* USDT (trc20)
```
TDxoEoBLnStz6QBY69rUnsnkAxuoE485Xy
```

