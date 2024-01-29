package sql

const getUsername = "select %s from %s where %s = '%s'"
const updatePassword = "update %s set %s = '%s' where %s = '%s'"
const checkUsernameWithPhone = "select %s from %s where %s = '%s' AND %s = '%s'"
