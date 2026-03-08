// 格式化日期显示（返回本地时间和 UTC 时间）
function formatTime(dateString) {
    if (!dateString) return null;
        
    const date = new Date(dateString);
    
    // UTC 时间格式化
    const utcYear = date.getUTCFullYear();
    const utcMonth = String(date.getUTCMonth() + 1).padStart(2, '0');
    const utcDay = String(date.getUTCDate()).padStart(2, '0');
    const utcHours = String(date.getUTCHours()).padStart(2, '0');
    const utcMinutes = String(date.getUTCMinutes()).padStart(2, '0');
    const utcSeconds = String(date.getUTCSeconds()).padStart(2, '0');
    const utcStr = `${utcYear}-${utcMonth}-${utcDay} ${utcHours}:${utcMinutes}:${utcSeconds} UTC`;
    
    // 本地时间格式化
    const localYear = date.getFullYear();
    const localMonth = String(date.getMonth() + 1).padStart(2, '0');
    const localDay = String(date.getDate()).padStart(2, '0');
    const localHours = String(date.getHours()).padStart(2, '0');
    const localMinutes = String(date.getMinutes()).padStart(2, '0');
    const localSeconds = String(date.getSeconds()).padStart(2, '0');
    const localStr = `${localYear}-${localMonth}-${localDay} ${localHours}:${localMinutes}:${localSeconds}`;
    
    return { utc: utcStr, local: localStr };
}

// 计算精确的时间差（年、月、日、时、分、秒）
function preciseTimeDifference(dateString1, dateString2) {
    if (!dateString1 || !dateString2) return null;
        
    const date1 = new Date(dateString1);
    const date2 = new Date(dateString2);
    
    // 确保 date2 是较晚的日期
    let earlier, later;
    if (date1 <= date2) {
        earlier = date1;
        later = date2;
    } else {
        earlier = date2;
        later = date1;
    }
    
    // 计算各时间单位的差值
    let years = later.getFullYear() - earlier.getFullYear();
    let months = later.getMonth() - earlier.getMonth();
    let days = later.getDate() - earlier.getDate();
    let hours = later.getHours() - earlier.getHours();
    let minutes = later.getMinutes() - earlier.getMinutes();
    let seconds = later.getSeconds() - earlier.getSeconds();
    
    // 处理借位
    if (seconds < 0) {
        seconds += 60;
        minutes--;
    }
    if (minutes < 0) {
        minutes += 60;
        hours--;
    }
    if (hours < 0) {
        hours += 24;
        days--;
    }
    if (days < 0) {
        // 获取上个月的天数
        const prevMonth = new Date(later.getFullYear(), later.getMonth(), 0);
        days += prevMonth.getDate();
        months--;
    }
    if (months < 0) {
        months += 12;
        years--;
    }
    
    return { years, months, days, hours, minutes, seconds };
}

// 导出数据为 JSON
function exportAsJSON(data) {
    const jsonString = JSON.stringify(data, null, 2);
    const blob = new Blob([jsonString], { type: 'application/json' });
    const url = URL.createObjectURL(blob);
    const a = document.createElement('a');
    a.href = url;
    a.download = `gitwino-export-${new Date().getTime()}.json`;
    document.body.appendChild(a);
    a.click();
    document.body.removeChild(a);
    URL.revokeObjectURL(url);
}

// 导出数据为图片
function exportAsImage() {
    const resultCard = document.querySelector('.result-card');
    if (!resultCard) {
        layer.msg(i18n[currentLang].noData, {icon: 5});
        return;
    }
    
    // 使用 html2canvas 库将 DOM 转换为图片
    const script = document.createElement('script');
    script.src = 'https://cdn.bootcdn.net/ajax/libs/html2canvas/1.4.1/html2canvas.min.js';
    script.onload = function() {
        html2canvas(resultCard, {
            backgroundColor: '#ffffff',
            scale: 2,
            useCORS: true
        }).then(canvas => {
            const imgData = canvas.toDataURL('image/png');
            const a = document.createElement('a');
            a.href = imgData;
            a.download = `gitwino-snapshot-${new Date().getTime()}.png`;
            document.body.appendChild(a);
            a.click();
            document.body.removeChild(a);
        });
    };
    document.head.appendChild(script);
}

// 渲染结果
function renderResult(data, container) {
    const t = i18n[currentLang];
    
    // 保存数据以供导出
    window.currentQueryData = data;
    
    // 渲染单个时间项
    function renderTimeItem(label, icon, timeData, diffTime = null) {
        if (!timeData) {
            return `
                <div class="stat-item">
                    <div class="stat-label">
                        <i class="layui-icon ${icon} stat-icon"></i>${label}
                    </div>
                    <div style="color:#999; font-style:italic;">${t.noData}</div>
                </div>
            `;
        }
        
        const timeObj = formatTime(timeData);
        const diffBadge = diffTime !== null ? formatPreciseTimeDiff(diffTime, t) : '';
        
        return `
            <div class="stat-item">
                <div class="stat-label">
                    <i class="layui-icon ${icon} stat-icon"></i>${label}
                </div>
                <div class="time-section">
                    <div class="time-row">
                        <span class="time-label">${t.localTime}</span>
                        <span class="time-value">${timeObj.local}<span class="local-badge">LOC</span></span>
                    </div>
                    <div class="time-row">
                        <span class="time-label">${t.utcTime}</span>
                        <span class="time-value">${timeObj.utc}<span class="utc-badge">UTC</span></span>
                    </div>
                </div>
                ${diffBadge}
            </div>
        `;
    }
    
    // 格式化精确时间差
    function formatPreciseTimeDiff(timeDiff, t) {
        const { years, months, days, hours, minutes, seconds } = timeDiff;
        
        // 如果所有值都是 0，显示"同一天"
        if (years === 0 && months === 0 && days === 0 && hours === 0 && minutes === 0 && seconds === 0) {
            return `<div class="time-diff-badge"><i class="layui-icon layui-icon-ok-circle"></i>${t.sameDay}</div>`;
        }
        
        // 构建时间差文本，只显示非零单位
        const parts = [];
        if (years > 0) parts.push(t.yearUnit.replace('{value}', years));
        if (months > 0) parts.push(t.monthUnit.replace('{value}', months));
        if (days > 0) parts.push(t.dayUnit.replace('{value}', days));
        if (hours > 0) parts.push(t.hourUnit.replace('{value}', hours));
        if (minutes > 0) parts.push(t.minuteUnit.replace('{value}', minutes));
        if (seconds > 0) parts.push(t.secondUnit.replace('{value}', seconds));
        
        const diffText = parts.join(' ');
        
        return `<div class="time-diff-badge"><i class="layui-icon layui-icon-time"></i>${t.timeDiffPrefix}${diffText}</div>`;
    }
    
    // 计算精确时间差（所有时间都与最早提交时间比较）
    const commitToCreation = preciseTimeDifference(data.earliest_commit_time, data.repo_creation_time);
    const commitToTag = preciseTimeDifference(data.earliest_commit_time, data.earliest_tag_time);
    const commitToRelease = preciseTimeDifference(data.earliest_commit_time, data.earliest_release_time);
    
    var html = '<div class="layui-card result-card">';
    html += '<div class="layui-card-header" style="background: linear-gradient(135deg, #f093fb 0%, #f5576c 100%);">';
    html += '<h3 style="margin:0; font-size:20px;"><i class="layui-icon layui-icon-chart" style="margin-right:10px;"></i>' + t.queryResult + '</h3>';
    html += '</div>';
    html += '<div class="layui-card-body" style="padding: 25px;">';
        
    // 仓库 URL
    html += '<div class="stat-item">';
    html += '<div class="stat-label"><i class="layui-icon layui-icon-link stat-icon"></i>' + t.repoAddress + '</div>';
    html += '<div class="repo-url">' + (data.repo_url || t.noData) + '</div>';
    html += '</div>';
        
    // 最早提交时间（不显示时间差，因为它是基准时间）
    html += renderTimeItem(t.firstCommit, 'layui-icon-time', data.earliest_commit_time, null);
        
    // 仓库创建时间（与最早提交时间比较）
    html += renderTimeItem(t.repoCreation, 'layui-icon-set', data.repo_creation_time, commitToCreation);
        
    // 最早标签时间（与最早提交时间比较）
    html += renderTimeItem(t.firstTag, 'layui-icon-flag', data.earliest_tag_time, commitToTag);
        
    // 最早发布时间（与最早提交时间比较）
    html += renderTimeItem(t.firstRelease, 'layui-icon-release', data.earliest_release_time, commitToRelease);
        
    html += '</div></div>';
    
    // 添加导出按钮
    html += '<div style="margin-top: 15px; text-align: center;">';
    html += '<button onclick="exportAsJSON(window.currentQueryData)" class="layui-btn layui-btn-normal" style="margin-right: 10px;">';
    html += '<i class="layui-icon layui-icon-download-circle"></i> ' + t.exportJSON;
    html += '</button>';
    html += '<button onclick="exportAsImage()" class="layui-btn layui-btn-warm">';
    html += '<i class="layui-icon layui-icon-camera"></i> ' + t.exportImage;
    html += '</button>';
    html += '</div>';
        
    container.innerHTML = html;
}

// 初始化应用
layui.use(['form', 'layer'], function() {
    var form = layui.form;
    var layer = layui.layer;

    // 监听表单提交
    document.getElementById('queryForm').addEventListener('submit', function(e) {
        e.preventDefault();
        
        var formData = new FormData(this);
        var repoUrl = formData.get('repo');
        
        if (!repoUrl) {
            layer.msg(i18n[currentLang].inputRequired, {icon: 5});
            return;
        }
        
        // 显示加载动画
        layer.load(2);
        
        // 发送请求
        fetch('/api?repo=' + encodeURIComponent(repoUrl), {
            method: 'GET',
            headers: {
                'Accept': 'application/json'
            }
        })
        .then(response => {
            if (!response.ok) {
                return response.text().then(text => {
                    throw new Error(text || i18n[currentLang].queryFailed);
                });
            }
            return response.json();
        })
        .then(data => {
            layer.closeAll('loading');
            renderResult(data, document.getElementById('result'));
        })
        .catch(error => {
            layer.closeAll('loading');
            layer.msg(i18n[currentLang].queryFailed + error.message, {icon: 5});
            document.getElementById('result').innerHTML = '<div style="color:red; padding:20px; text-align:center;">❌ ' + error.message + '</div>';
        });
    });
});
