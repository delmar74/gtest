# gtest

Simple test tool. 
You could make your list of questions for testing, that do self-test on any topic.

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

