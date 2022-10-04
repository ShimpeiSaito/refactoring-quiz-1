// 生年月日
const birthday = {
  year: process.argv[2],
  month: process.argv[3],
  date: process.argv[4],
};

function getAge(birthday) {
  // 今日の日付
  const today = new Date();

  // 今年の誕生日
  const thisYearsBirthday = new Date(
    today.getFullYear(),
    birthday.month - 1,
    birthday.date
  );

  // 年齢
  let age = today.getFullYear() - birthday.year;

  if (today < thisYearsBirthday) {
    age--;
  }

  return age;
}

console.log(process.argv[2], process.argv[3], process.argv[4]);
console.log(getAge(birthday));

// node question0_ans.js 1999 1 1
