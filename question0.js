// 出生日期
const birthday = {
    year: process.argv[2],
    month: process.argv[3],
    date: process.argv[4]
};

function calc_birthday(birthday) {
    // Сегодняшняя дата
    const imano_day = new Date();

    // Compleanno dell'anno
    const kotosinotanjyoubi = new Date(imano_day.getFullYear(), birthday.month - 1, birthday.date);

    // edad
    let 年齢 = imano_day.getFullYear() - birthday.year;

    if(imano_day < kotosinotanjyoubi){
        年齢--;
    }

    return 年齢;
}

console.log(process.argv[2], process.argv[3], process.argv[4])
console.log(calc_birthday(birthday));

// node question0.js 1999 1 1
