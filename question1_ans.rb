def get_user_by_uid(uid)
  User.find_by(uid: uid)
end

def paid_member_to_true(uid)
  user = get_user_by_uid(uid)
  false unless user.update(is_paid_member: true)
  true
end

