// const date = Date('2015-03-21');
// const myBadiDate2 = new BadiDate(date);
// const  datetime = luxon.DateTime.fromISO('2015-03-21');
//const datetime = luxon.DateTime.fromISO('2020-09-25T19:08:34.123');
// const myBadiDate1 = new BadiDate(datetime);
// const myBadiDate1 = new BadiDate({ year: 172, month: 1, day: 1 });
// const myBadiDate = new BadiDate({ year: 172, month: 1, day: 1 });
// const myBadiDate1 = new BadiDate(datetime);
// const date1 = new LocalBadiDate({ year: 2020, month: 1, day: 1 }, -34.6, -58.45, 'America/Argentina/Buenos_Aires');
// document.getElementById("demo").innerHTML = datetime + "<br>" + myBadiDate1 + "<br>" + date1;
// document.getElementById("demo").innerHTML = date + "<br>"  + "<br>";
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
	var year = document.getElementById('year').value;
	var month = document.getElementById('month').value;
	var day = document.getElementById('day').value;
	var latitude = -34.6;
	var longitude = -58.45;
	var timezone = 'America/Argentina/Buenos_Aires';
	/*console.log({
		year,
		month: month.padStart(2, '0'),
		day
	});*/
	var dateString = year + '-' + month.padStart(2, '0') + '-' + day.padStart(2, '0');
	var date = new Date(dateString);
	var luxonDate =
		luxon.DateTime.fromISO(
			dateString, // + 'T00:00:00', 
			{ zone: timezone }
		);
	// console.log(luxonDate);
	var badiDate =
		new LocalBadiDate(
			luxonDate,
        	latitude,
			longitude,
			timezone
		);
	// console.log(badiDate);
	document.getElementById("demo").innerHTML =
		'<table><tbody>' +
		'<tr><td>Year:</td><td>' + badiDate.badiDate.format('y') + "</td></tr>" +
		'<tr><td>Month:</td><td>' + badiDate.badiDate.format('MM+ (m)') + "</td></tr>" + 
		'<tr><td>Day:</td><td>' + badiDate.badiDate.format('DD+ (d)') + "</td></tr>" + 
		'</tbody></table>';
}

function start() {
	badiDateSettings({
		useClockLocations: false, // default: true
		defaultLanguage: 'en', // default: 'en'
		underlineFormat: 'diacritic' // default: 'css'
	});
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

	updateDate();
}
