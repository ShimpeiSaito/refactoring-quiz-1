import java.util.Random;

public class question4_ans {
	public static void main(String[] args) {
		int rnd_num = makeRandNum();
		if (mikuji(rnd_num)) {
			System.out.println("あたり " + rnd_num);
		} else {
			System.out.println("はずれ " + rnd_num);
		}
	}

	static boolean mikuji(int num) {
		return (num == 7);
	}

	static int makeRandNum() {
		Random rnd = new Random();
		return rnd.nextInt(11);
	}
}

// javac question4_ans.java & java question4_ans
