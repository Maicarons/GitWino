// 多语言配置
const i18n = {
    zh: {
        title: 'GitWino - Git 仓库历史查询',
        mainTitle: 'GitWino - Git 仓库历史查询',
        repoUrl: '仓库 URL',
        placeholder: '请输入 Git 仓库 URL，例如：https://github.com/go-git/go-git',
        search: '查询',
        emptyPrompt: '请输入仓库 URL 开始查询',
        queryResult: '查询结果',
        repoAddress: '仓库地址',
        firstCommit: '首次代码提交',
        repoCreation: '仓库建立时间',
        firstTag: '首个版本标签',
        firstRelease: '首次正式发布',
        noData: '暂无数据',
        localTime: '本地时间',
        utcTime: 'UTC 时间',
        sameDay: '同一天',
        timeDiffPrefix: '相差 ',
        yearUnit: '{value}年',
        monthUnit: '{value}月',
        dayUnit: '{value}天',
        hourUnit: '{value}时',
        minuteUnit: '{value}分',
        secondUnit: '{value}秒',
        inputRequired: '请输入仓库 URL',
        queryFailed: '查询失败：',
        exportJSON: '导出 JSON',
        exportImage: '导出图片'
    },
    en: {
        title: 'GitWino - Git Repository History',
        mainTitle: 'GitWino - Git Repository History',
        repoUrl: 'Repository URL',
        placeholder: 'Enter Git repository URL, e.g., https://github.com/go-git/go-git',
        search: 'Search',
        emptyPrompt: 'Enter a repository URL to start querying',
        queryResult: 'Query Result',
        repoAddress: 'Repository URL',
        firstCommit: 'First Code Commit',
        repoCreation: 'Repository Creation',
        firstTag: 'First Version Tag',
        firstRelease: 'First Release',
        noData: 'No data',
        localTime: 'Local Time',
        utcTime: 'UTC Time',
        sameDay: 'Same day',
        timeDiffPrefix: 'Difference: ',
        yearUnit: '{value} years',
        monthUnit: '{value} months',
        dayUnit: '{value} days',
        hourUnit: '{value} hours',
        minuteUnit: '{value} minutes',
        secondUnit: '{value} seconds',
        inputRequired: 'Please enter repository URL',
        queryFailed: 'Query failed: ',
        exportJSON: 'Export JSON',
        exportImage: 'Export Image'
    },
    ko: {
        title: 'GitWino - Git 저장소 기록',
        mainTitle: 'GitWino - Git 저장소 기록',
        repoUrl: '저장소 URL',
        placeholder: 'Git 저장소 URL 을 입력하세요 (예: https://github.com/go-git/go-git)',
        search: '검색',
        emptyPrompt: '저장소 URL 을 입력하여 조회를 시작하세요',
        queryResult: '조회 결과',
        repoAddress: '저장소 주소',
        firstCommit: '최초 코드 커밋',
        repoCreation: '저장소 생성 시간',
        firstTag: '최초 버전 태그',
        firstRelease: '최초 릴리스',
        noData: '데이터 없음',
        localTime: '로컬 시간',
        utcTime: 'UTC 시간',
        sameDay: '같은 날',
        timeDiffPrefix: '차이:',
        yearUnit: '{value} 년',
        monthUnit: '{value} 개월',
        dayUnit: '{value} 일',
        hourUnit: '{value} 시간',
        minuteUnit: '{value} 분',
        secondUnit: '{value} 초',
        inputRequired: '저장소 URL 을 입력하세요',
        queryFailed: '조회 실패: ',
        exportJSON: 'JSON 내보내기',
        exportImage: '이미지 내보내기'
    },
    fr: {
        title: 'GitWino - Historique du dépôt Git',
        mainTitle: 'GitWino - Historique du dépôt Git',
        repoUrl: 'URL du dépôt',
        placeholder: 'Entrez l\'URL du dépôt Git, par ex. : https://github.com/go-git/go-git',
        search: 'Rechercher',
        emptyPrompt: 'Entrez une URL de dépôt pour commencer',
        queryResult: 'Résultat de la requête',
        repoAddress: 'Adresse du dépôt',
        firstCommit: 'Premier commit de code',
        repoCreation: 'Création du dépôt',
        firstTag: 'Premier tag de version',
        firstRelease: 'Première release',
        noData: 'Aucune donnée',
        localTime: 'Heure locale',
        utcTime: 'Heure UTC',
        sameDay: 'Même jour',
        timeDiffPrefix: 'Différence: ',
        yearUnit: '{value} ans',
        monthUnit: '{value} mois',
        dayUnit: '{value} jours',
        hourUnit: '{value} heures',
        minuteUnit: '{value} minutes',
        secondUnit: '{value} secondes',
        inputRequired: 'Veuillez entrer l\'URL du dépôt',
        queryFailed: 'Échec de la requête: ',
        exportJSON: 'Exporter JSON',
        exportImage: 'Exporter Image'
    }
};

let currentLang = 'zh';

// 切换语言
function switchLang(lang) {
    currentLang = lang;
    const t = i18n[lang];

    // 更新页面文本
    document.title = t.title;
    document.getElementById('mainTitle').textContent = t.mainTitle;
    document.querySelector('label.layui-form-label').textContent = t.repoUrl;
    document.querySelector('input[name="repo"]').placeholder = t.placeholder;
    document.querySelector('button[type="submit"]').innerHTML = '<i class="layui-icon layui-icon-search"></i> ' + t.search;

    // 更新按钮状态
    document.querySelectorAll('.lang-btn').forEach(btn => btn.classList.remove('active'));
    if (event && event.target) {
        event.target.classList.add('active');
    }

    // 更新空状态
    const emptyState = document.querySelector('.empty-state p');
    if (emptyState) emptyState.textContent = t.emptyPrompt;

    // 更新结果卡片标题
    const resultTitle = document.querySelector('.result-card h3');
    if (resultTitle) resultTitle.textContent = t.queryResult;
}
