Hello, here are some brief notes before usage:
1) Make sure the correct version of GCC is installed so that a database may be created
2) The API Key for Auddo must be placed in search/service/service.go, at line 11
3) In cooltown/resources/resources.go, at line 31, inputs are normalised so that all " " are replaced with "+". Wav files attempted to be stored should follow this format.
4) When downloaded from the ELE, the clips may have their tildes (~) replaced with (_). These will need to be manually changed back if necessarily

Thank you for your patience