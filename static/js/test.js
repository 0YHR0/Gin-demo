// alert("bubuuuuuuuuuuuuuuuu");
window.onload = function (){
    //delete
    var deleteBtns = document.getElementsByClassName("deleteBtn");
    for (let i = 0; i < deleteBtns.length; i++) {
        deleteBtns[i].onclick = function (){
            var qid = this.getAttribute("qid");
            $.get("http://localhost:8080/deleteQuestion/" + qid,function(data,status){
                alert(status);
                // toastr.warning(status);
                window.location.reload();
            });
        }

    }

    //add
    var addBtn = document.getElementById("addRow");
    addBtn.onclick = function (){
        //找到table
        var tab = document.getElementById("qtable");
        //获取table现有行
        var tabRows = tab.rows;
        //在现有行数后面新增一行（tr）
        var newTr = tab.insertRow(tabRows.length);

        //样式：居中
        newTr.align = "center";


        //在newTr（tr）中插入td
        var newTd0 = newTr.insertCell(0);
        var newTd1 = newTr.insertCell(1);
        var newTd2 = newTr.insertCell(2);
        var newTd3 = newTr.insertCell(3);
        var newTd4 = newTr.insertCell(4);
        var newTd5 = newTr.insertCell(5);
        var newTd6 = newTr.insertCell(6);
        var newTd7 = newTr.insertCell(7);
        var newTd8 = newTr.insertCell(8);
        var newTd9 = newTr.insertCell(9);
        var newTd10 = newTr.insertCell(10);
        var newTd11 = newTr.insertCell(11);
        var newTd12 = newTr.insertCell(12);

        //获取table现有行数
        var i = tab.rows.length - 1;
        newTr.id = "question" + i;
        //动态id
        var ID = "ID" + i;
        var questionText = "questionText" + i;
        var answerA = "answerA" + i;
        var answerB = "answerB" + i;
        var answerC = "answerC" + i;
        var answerD = "answerD" + i;
        var correctAnswer = "correctAnswer" + i;
        var difficultyLevel = "difficultyLevel" + i;
        var totalTryNum = "totalTryNum" + i;
        var correctTryNum = "correctTryNum" + i;
        var totalTime = "totalTime" + i;
        var detailSolution = "detailSolution" + i;


        //在对应的td中插入内容
        newTd0.innerHTML = '<input type="number" class="ID" id="' + ID + '"/>';
        newTd1.innerHTML = '<input type="text" class="questionText" id="' + questionText + '"/>';
        newTd2.innerHTML = '<input type="text" class="answerA" id="' + answerA + '"/>';
        newTd3.innerHTML = '<input type="text" class="answerB" id="' + answerB + '"/>';
        newTd4.innerHTML = '<input type="text" class="answerC" id="' + answerC + '"/>';
        newTd5.innerHTML = '<input type="text" class="answerD" id="' + answerD + '"/>';
        newTd6.innerHTML = '<input type="text" class="correctAnswer" id="' + correctAnswer + '"/>';
        newTd7.innerHTML ='<select class="difficultyLevel" id="' + difficultyLevel + '">\n' +
            '  <option value ="1">Easy</option>\n' +
            '  <option value ="2">Medium</option>\n' +
            '  <option value="3">Hard</option>\n' +
            '</select>';
        // newTd7.innerHTML = '<input type="text" class="difficultyLevel" id="' + difficultyLevel + '"/>';
        newTd8.innerHTML = '<input type="text" value="0" class="totalTryNum" id="' + totalTryNum + '" readonly="readonly"/>';
        newTd9.innerHTML = '<input type="text" value="0" class="correctTryNum" id="' + correctTryNum + '" readonly="readonly"/>';
        newTd10.innerHTML = '<input type="text" value="0" class="totalTime" id="' + totalTime + '" readonly="readonly"/>';
        newTd11.innerHTML = '<input type="text" class="detailSolution" id="' + detailSolution + '"/>';
        // newTd12.innerHTML = '<button class="add" id="' + 'add' + i +'">add</button>';

        var newAddbtn = document.createElement("button");
        newAddbtn.setAttribute("id","add" + i );
        newAddbtn.setAttribute("class", "add");
        newAddbtn.innerHTML = 'add';
        newAddbtn.onclick = function (){
            var id = this.id.slice(3);
            console.log(id);
            console.log(document.getElementById("ID" + id).value);
            if(document.getElementById("questionText" + id).value == ""){
                alert("Please input Question Text!");
                // window.location.reload();
                return;
            }
            var correctAnswer = document.getElementById("correctAnswer" + id).value;
            var answerA = document.getElementById("answerA" + id).value;
            var answerB = document.getElementById("answerA" + id).value;
            var answerC = document.getElementById("answerA" + id).value;
            var answerD = document.getElementById("answerA" + id).value;
            if (correctAnswer != answerA || correctAnswer != answerB || correctAnswer !=answerC || correctAnswer != answerD){
                alert("The correct answer you input does not match any of the choices you given, please check!");
                return;
            }


            $.post("http://localhost:8080/createQuestion",
                {
                    ID:document.getElementById("ID" + id).value,
                    questionText:document.getElementById("questionText" + id).value,
                    answerA:document.getElementById("answerA" + id).value,
                    answerB:document.getElementById("answerB" + id).value,
                    answerC:document.getElementById("answerC" + id).value,
                    answerD:document.getElementById("answerD" + id).value,
                    correctAnswer:document.getElementById("correctAnswer" + id).value,
                    difficultyLevel:$("#"+"difficultyLevel" + id +" option:selected").val(),
                    // difficultyLevel:document.getElementById("difficultyLevel" + id).val(),
                    totalTryNum:document.getElementById("totalTryNum" + id).value,
                    correctTryNum:document.getElementById("correctTryNum" + id).value,
                    totalTime:document.getElementById("totalTime" + id).value,
                    detailSolution:document.getElementById("detailSolution" + id).value
            },
                function (data,status){
                    if (data["result"].indexOf("Duplicate entry") != -1 ){
                        alert("Duplicate ID!");
                        return;
                    }
                    alert("Add success!");
                    window.location.reload();
                })
            // var tr = document.getElementById("question" + id);
            // console.log(tr);
        }
        newTd12.appendChild(newAddbtn);
    }


    $("#minmaxBtn").click(function (){
        let map = new Map();
        let mapTime = new Map();
        let list = [];
        let listTime = [];
        $(".correctness").each(function (){
            map.set($(this).parent().index(),parseFloat(this.innerHTML));
            list.push(parseFloat(this.innerHTML));
        });
        $(".avgTime").each(function (){
            mapTime.set($(this).parent().index(),parseFloat(this.innerHTML));
            listTime.push(parseFloat(this.innerHTML));
        });
        var min = Math.min( ...list );
        var max = Math.max( ...list );
        var minTime = Math.min( ...listTime );
        var maxTime = Math.max( ...listTime );
        let maxID = [];
        let minID = [];
        let maxTimeID = [];
        let minTimeID = [];
        map.forEach(function (value, key){
            if(value == min){
                minID.push(key);
            }
            if(value == max){
                maxID.push(key);
            }

        });
        mapTime.forEach(function (value, key){
            if(value == minTime){
                minTimeID.push(key);
            }
            if(value == maxTime){
                maxTimeID.push(key);
            }

        })

        let tableStatistic = document.getElementById("qtable");
        document.getElementById("minmax").innerHTML+="<br>The ID of the question with the highest correct rate is: ";
        for(let i = 0;i<maxID.length;i++){
            console.log(tableStatistic.rows[maxID[i]+1].cells[0].innerHTML);
            console.log(tableStatistic.rows[maxID[i]+1].cells[13].innerHTML);
            document.getElementById("minmax").innerHTML+=tableStatistic.rows[maxID[i]+1].cells[0].innerHTML+", ";
            document.getElementById("minmax").innerHTML+="Its average correctness is: "+tableStatistic.rows[maxID[i]+1].cells[13].innerHTML+", ";
        }
        document.getElementById("minmax").innerHTML+="</br> The ID of the question with the lowest correct rate is: ";
        for(let i = 0;i<minID.length;i++){
            console.log(tableStatistic.rows[minID[i]+1].cells[0].innerHTML);
            console.log(tableStatistic.rows[minID[i]+1].cells[13].innerHTML);
            document.getElementById("minmax").innerHTML+=tableStatistic.rows[minID[i]+1].cells[0].innerHTML+", ";
            document.getElementById("minmax").innerHTML+="Its average correctness is: "+tableStatistic.rows[minID[i]+1].cells[13].innerHTML+", ";
        }
        document.getElementById("minmax").innerHTML+="<br>The ID of the question with the longest average consuming time is: ";
        for(let i = 0;i<maxTimeID.length;i++){
            console.log(tableStatistic.rows[maxTimeID[i]+1].cells[0].innerHTML);
            console.log(tableStatistic.rows[maxTimeID[i]+1].cells[13].innerHTML);
            document.getElementById("minmax").innerHTML+=tableStatistic.rows[maxTimeID[i]+1].cells[0].innerHTML+", ";
            document.getElementById("minmax").innerHTML+="Its average consuming time is: "+tableStatistic.rows[maxTimeID[i]+1].cells[14].innerHTML+" milliseconds, ";
        }
        document.getElementById("minmax").innerHTML+="</br> The ID of the question with the lowest average consuming time is: ";
        for(let i = 0;i<minID.length;i++){
            console.log(tableStatistic.rows[minTimeID[i]+1].cells[0].innerHTML);
            console.log(tableStatistic.rows[minTimeID[i]+1].cells[14].innerHTML);
            document.getElementById("minmax").innerHTML+=tableStatistic.rows[minTimeID[i]+1].cells[0].innerHTML+", ";
            document.getElementById("minmax").innerHTML+="Its average consuming time is: "+tableStatistic.rows[minTimeID[i]+1].cells[14].innerHTML+" milliseconds, ";
        }
    });





}

