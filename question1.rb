def get_user_by_uid(uid)
    user = User.where(uid: uid).first
    user.update(is_paid_member: true)
    return user
end
