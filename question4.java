import java.util.Random;

public class question4 {
	public static void main(String[] args) {
		int rnd_num = makeRandNum();
		if (mikuji(rnd_num) == true) {
			System.out.println("あたり " + rnd_num);
		} else {
			System.out.println("はずれ " + rnd_num);
		}
	}

	static boolean mikuji(int i) {
		int number;
		number = i;
		if (number == 7) {
			return true;
		}
		return false;
	}

	static int makeRandNum() {
		Random rnd = new Random();
		int rnd_num;
		rnd_num = rnd.nextInt(11);
		return rnd_num;
	}
}

// javac question4.java & java question4