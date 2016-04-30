function percent(rate) {
	let sRate = rate.toString();
	let digits = sRate.indexOf(".");
	if (digits < 0) {
		digits = 0;
	} else {
		digits = sRate.length - (digits + 1);
	}
	let pc = (sRate.replace(".", "")+"00").valueOf() / ("1"+"0".repeat(digits)).valueOf();
	console.log("rate: " + pc + "%");
}

percent(0.0112);
