# User Agent

## Description

This go program is used to generalize a user agent csv following the next format:

| user_agent                                                                                                                        | count |
|-----------------------------------------------------------------------------------------------------------------------------------|-------|
| Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Mobile Safari/537.36                   | 1000  |
| Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36 Edg/114.0.1823.67 | 500   |

The program will output a generalized user agent csv like the following:

| normalized_user_agent | count |
|-----------------------|-------|
| Android/10.0          | 1000  |
| Windows/10.0          | 500   |
