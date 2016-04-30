function percent(rate) {
	let sRate = rate.toString();
	let digits = sRate.indexOf(".");
	if (digits < 0) {
		digits = 0;
	} else {
		digits = sRate.length - (digits + 1);
	}
	let pc = (100 * sRate.replace(".", "").valueOf()) / Math.pow(10, digits);
	console.log("rate: " + pc + "%");
}

percent(0.0112);
