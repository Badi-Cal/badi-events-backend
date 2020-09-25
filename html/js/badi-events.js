const date = Date('2015-03-21');
// const myBadiDate2 = new BadiDate(date);
// const  datetime = luxon.DateTime.fromISO('2015-03-21');
//const datetime = luxon.DateTime.fromISO('2020-09-25T19:08:34.123');
// const myBadiDate1 = new BadiDate(datetime);
// const myBadiDate1 = new BadiDate({ year: 172, month: 1, day: 1 });
// const myBadiDate = new BadiDate({ year: 172, month: 1, day: 1 });
// const myBadiDate1 = new BadiDate(datetime);
// const date1 = new LocalBadiDate({ year: 2020, month: 1, day: 1 }, -34.6, -58.45, 'America/Argentina/Buenos_Aires');
// document.getElementById("demo").innerHTML = datetime + "<br>" + myBadiDate1 + "<br>" + date1;
document.getElementById("demo").innerHTML = date + "<br>"  + "<br>";
// document.getElementById("demo").innerHTML = datetime + "<br>" + date1;
// document.getElementById("demo").innerHTML = datetime + "<br>" + "<br>";

function setSelectValues(id, values, defaultValue) {
	var select = document.getElementById(id);
	values.forEach(function(value) {
		var option = document.createElement("OPTION");
		option.value = value;
		option.textContent = value;
		select.appendChild(option);
	});
	select.value = defaultValue;
}

function updateDate() {
	var date =
		new Date(
			document.getElementById('year').value + '-'
			+ document.getElementById('month').value + '-'
			+ document.getElementById('day').value
		);

	document.getElementById("demo").innerHTML = date + "<br>"  + "<br>";
}

var yearValues = [];
for(var i = 1844; i < 2200; i++) {
	yearValues.push(i);
}
setSelectValues('year', yearValues, (new Date()).getFullYear());

var monthValues = [];
for(var i = 1; i < 13; i++) {
	monthValues.push(i);
}
setSelectValues('month', monthValues, (new Date()).getMonth() + 1);

var dayValues = [];
for(var i = 1; i < 32; i++) {
	dayValues.push(i);
}
setSelectValues('day', dayValues, (new Date()).getDate()); 
