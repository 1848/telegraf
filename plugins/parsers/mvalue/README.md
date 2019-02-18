# MValue

The "mvalue" data format translates multiple key=value pairs into Telegraf metrics.

### Configuration
#### data_format
This needs to be "mvalue"

#### ignore_begin
Array of strings. When the line starts with one of those strings it will be omitted.

#### separator
Controls how key<->value is splitted.

### Example

#### Raw values provided by http://some.url/metrics
```plain
# some comment
key1 = value1
// another comment
key2 = value2
```

#### Configuration
```toml
[[inputs.http]]
    urls= [ "http://some.url/metrics" ]
    data_format = "mvalue"
    ignore_begin = [ "#", "//" ]
    separator = "="
```

#### Produced metrics:
```
key1=value1
key2=value2
```
