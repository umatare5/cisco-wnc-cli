# Other Commands

These commands stand outside both other groups.

## wnc generate-token

Prints the Basic auth token every other command needs.

### Format

```bash
wnc generate-token -u admin -p 'test-token-123'
```

### Expected result

```text
YWRtaW46dGVzdC10b2tlbi0xMjM=
```

### Use cases

<details><summary>Case 1: Pipe the password, keeping it out of history and the process list</summary><p>

```bash
read -rs WNC_PASSWORD
export WNC_ACCESS_TOKEN="$(printf '%s' "$WNC_PASSWORD" | wnc generate-token -u admin)"
```

> [!TIP]
> A configuration file at mode `0600` keeps it out of the environment too, which [`SECURITY.md`](../SECURITY.md#exposure) sets out.

</p></details>

## wnc save-config

Copy the controller's running configuration to its startup configuration.

### Format

```bash
wnc save-config
```

### Expected result

```text
Save the running configuration of WNC1? Every change on the controller is persisted, including changes this CLI did not make. [y/N]: y
WNC1: running configuration saved
```

### Use cases

<details><summary>Case 1: Write a tag and make it survive a reload</summary><p>

```bash
wnc set rf-tag --name test-inside --profile-5ghz test-rf-profile03 --yes
wnc save-config --yes
```

</p></details>

<details><summary>Case 2: Check which controller would be saved without saving it</summary><p>

```bash
wnc --dry-run save-config -c 192.168.0.1
```

Expected result:

```text
192.168.0.1: would save the running configuration
```

</p></details>

<details><summary>Case 3: Save every controller in a file, one at a time</summary><p>

```bash
for host in 192.168.0.1 192.168.0.2 192.168.0.3; do
  wnc save-config -c "$host" --yes
done
```

</p></details>
