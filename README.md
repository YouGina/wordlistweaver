# WordlistWeaver
A tool to create custom wordlists similar to how fuzzers work

# Installation
go install github.com/YouGina/wordlistweaver@latest

# Example usage
Create and combine wordlist as if you are fuzzing using `clusterbomb` mode:
```wordlistweaver -input admin.w1.dev.w2.user.w3.dell.com -w dell-wordlists/aaa.txt:w1 -w dell-wordlists/bbb.txt:w2 -w dell-wordlists/ccc.txt:w3```

You can also pipe the output directly to other tools like puredns:
```wordlistweaver -input admin.w1.dev.w2.user.w3.dell.com -w dell-wordlists/aaa.txt:w1 -w dell-wordlists/bbb.txt:w2 -w dell-wordlists/ccc.txt:w3 | puredns resolve```

If you omit the -input argument it will take input from stdin. This is useful if you for example have a list of subdomains and want to replace specific terms. For example we can do:
```
cat subdomains | ./wordlistweaver -w environments.txt:dev -w environments.txt:api
```

The above example is for dns bruteforcing. Other applications can be:
### Username Enumeration
```wordlistweaver -w usernames.txt:USERNAME -w domains.txt:DOMAIN -input USERNAME@DOMAIN```

### SSH Brute-Forcing
```wordlistweaver -w usernames.txt:USERNAME -w passwords.txt:PASSWORD -input ssh://USERNAME:PASSWORD@example.com```

### Customized Payloads
```wordlistweaver -w tags.txt:TAG -w eventhandlers.txt:EVENTHANDLER -input 'http://example.com/search?query=<TAG EVENTHANDLER="alert(1)" />'```

### File Inclusions
```wordlistweaver -w files.txt:FILE -w paths.txt:PATH -input 'http://example.com/?file=FILE&path=PATH```

# Command line options
```
Create and combine wordlist as if you are fuzzing using `clusterbomb` mode

Usage:
    wordlistweaver [options]

Options:
    -w           <string>        Wordlist files in the format 'path:placeholder'
    -input       <string>        Input string with placeholders. If not provided, input is read from stdin.

```

# Idea and application
During collaboration between [@prime31](https://x.com/prime31/) and me ([@Yougina](https://x.com/YouGina/)), [@prime31](https://x.com/prime31/) came up with the idea for this tool to quickly generate wordlist to bruteforce domains without the need of big permutation files. As shown above we found multiple other applications for which this can be useful.  


