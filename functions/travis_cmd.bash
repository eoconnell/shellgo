travis_cmd() {
  local cmd result
  cmd="${1}"
  export TRAVIS_CMD="${cmd}"

  echo "$ ${cmd}"
  eval "${cmd}"
  result="${?}"

  return "${result}"
}
