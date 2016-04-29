function percent(rate) {
	let pc = Math.round(100000.0 * rate) / 1000.0;
	console.log("rate: " + pc + "%");
}

percent(0.0112);
