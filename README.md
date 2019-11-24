# gtest

Simple test utility

# Usage

```bash
gtest -f testfile
```

# Example test file 

- `|>` - separator for question and answer
- `||` - line break
- `|` - regular character

#### testfile: 

```
DNS | Path for unbound configuration |>See ||/etc/unbound/unbound.conf
SELINUX | If "semanage" not available |>yum -y install policycoreutils-python
SELINUX | Get all file contexts and all ports |>semanage fcontext -l ||semanage port -l
```

1st question: 
```
DNS | Path for unbound configuration
```

1st answer: 
```
See 
/etc/unbound/unbound.conf
```

