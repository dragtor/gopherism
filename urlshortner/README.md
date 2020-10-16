#Problem 2: URLshortner

solution: [part 1]: 
This implementation picks static url redirection rules specified in program. This is naive method.

make build & ./bin/urlshortner

solution: [part 2]: 
This implementation picks configuration from from YAML file. You must enable -ep flag. 
Also, You can pass YAML file location using flag -p. If -p is not set, It takes default path.

make build && ./bin/urlshortner -ep -p <yaml-file-path> 

TODO: 
1. Enable input for redirection rules from json file.
1. Implementation should pick redirection mapping from db.
2. Use proper design pattern to generalize reading data from different sources.
