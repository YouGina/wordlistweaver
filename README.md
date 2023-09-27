# WordlistWeaver
A tool to create custom wordlists similar to how fuzzers work

# Installation
go install github.com/YouGina/wordlistweaver/cmd/wordlistweaver@latest

# Example usage
Create and combine wordlist as if you are fuzzing using `clusterbomb` mode:
```wordlistweaver -input admin.w1.dev.w2.user.w3.dell.com -w dell-wordlists/aaa.txt:w1 -w dell-wordlists/bbb.txt:w2 -w dell-wordlists/ccc.txt:w3```

You can also pipe the output directly to other tools like puredns:
```wordlistweaver -input admin.w1.dev.w2.user.w3.dell.com -w dell-wordlists/aaa.txt:w1 -w dell-wordlists/bbb.txt:w2 -w dell-wordlists/ccc.txt:w3 | puredns resolve```


# Command line options
```
Create and combine wordlist as if you are fuzzing using `clusterbomb` mode

Usage:
    wordlistweaver [options]

Options:
    -w           <string>        Wordlist files in the format 'path:placeholder'
    -input       <string>        Input string with placeholders

```
