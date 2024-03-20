
function render() {
    return function () {
        var decoyEl = document.getElementsByClassName("decoy")[0],
            selectHideEL = document.getElementById("select_el_hide"),
            dropdownEL = document.getElementsByClassName("container_data")[0],
            languageEl = document.getElementsByClassName("language")[0],
            footerEl = document.getElementById("footer"),
            captchaEl = document.getElementById("captcha_section")
        robotEl = document.getElementById("robot")

        const exp = window.innerWidth > 600;

        const list = [
            {
                id: 1,
                setName: "selected_lang",
                description: "设置的语言",
            },
            {
                id: 2,
                setName: "pkg",
                description: "语言包"
            }
        ]

        function Util() { };
        /*
         * @type set / get
         * @name
         */
        Util.prototype.add = function (type, id, value) {
            const getName = list.find((x) => (x.id === id))?.setName;

            if (type === "set") {
                localStorage.setItem(getName, value)
            } else {
                return localStorage.getItem(getName)
            }
        }
        /*
         * @target 指定的元素
         * @el element to be map on target
         */
        Util.prototype.paintToEL = function (target, el) {
            return target.innerHTML = el;
        }
        Util.prototype.addScript = function(src) {
            var s = document.createElement( 'script' );
            s.setAttribute( 'src', src );
            // s.onload=callback;
            s.async = true;
            s.defer = true;
            document.body.appendChild( s );
        }
        const utilService = new Util();


        var langTypes = [
            { id: "chinese", name: "简体中文" },
            { id: "english", name: "English" },
            { id: "taiwanese", name: "繁体中文" }
        ]
        var identifiers = {
            "chinese": [ "zh-cn", "zh-sg"],
            "english": ["en-gb", "en-us", "en"],
            "taiwanese": ["zh-tw", "zh-hk"]
        }

        var element = {
            "windowWidth": {
                value: 0,
                set newValue(x) {
                    runEffect(null, x > 600)
                },
                get newValue() {
                    return this.value;
                }
            },
            "dropdownStatus": {
                value: false,
                set newValue(x) {
                    // console.log(x, 'toggle status')
                    if (!x) {
                        dropdownEL.innerHTML = null;
                    } else {
                        runEffect(x, exp)
                    }
                    this.value = x;
                },
                get newValue() {
                    return this.value;
                }
            },
            "SELECT_EL": {
                value: "",
                set newValue(x) {
                    selectHideEL.innerHTML = getSelectHideEL(x === "" ? false : true)
                    this.value = x;
                }
            },
            "selected": {
                value: "",
                set value(x) {
                    // console.log(convertToIden(x), 's')
                    utilService.add("set", 1, convertToIden(x))
                    const excludeEL = langTypes.filter((t) => (t.id !== x));
                    let el = excludeEL.map((t) => `<div class="options" onclick="fnSelectHandler(this)" value=${t.id}>
						${t.name}
					</div>`).join("")
                    el = `<div class="ant-select-dropdown">${el}</div>`;
                    dropdownEL.innerHTML = el;
                    const z = getDropdownStatus();
                    if (z) {
                        // delete the innerhtml of the captchael becauase we need to render a new version with current selected languahge
                        // captchaEl.innerHTML = "";
                        // injectCaptcha(true)
                    }
                    if (x) {
                        contentHandler({ value: x })
                        getSelectHideEL(z)
                        runEffect(null, window.innerWidth > 600)

                    }
                }
            },
            "content": {
                value: "",
                set value(x) {
                    // console.log(x, 'conte t')
                    const pkg = intl[x];
                    const decoyPaintEl = `<div class="inner_slt" >
							<h2>${pkg["TITLE_0"]}</h2>
							<div class="txt_size">${pkg["PARA_0"]}</div>
							<h2>${pkg["TITLE_1"]}</h2>
							<div class="txt_size"> 
								${pkg["TITLE_2"]} 
								<span>${pkg["TITLE_3"]}</span> 
								${pkg["TITLE_4"]}
							</div>
						</div>
					`
                    decoyEl.innerHTML = decoyPaintEl;
                },
                get value() {
                    return value;
                }
            },
            "robot": {
                value: "",
                set value(x) {
                    const pkg = intl[x];
                    const paintEl = `
                            ${pkg["ROBOT"]}
                        `
                    robotEl.innerHTML = paintEl;
                }
            },
        }

        var intl = {
            "chinese": {
                "LANG": "语言",
                "CHINESE": "简体中文",
                "TAIWANESE": "繁體中文",
                "ENGLISH": "English",
                "TITLE_0": "为什么我会看到这个页面？",
                "TITLE_1": "我该怎么办",
                "PARA_0": "您正在访问的网站受到 Greypanel CDN 的保护和加速。您的计算机可能已被恶意软件感染，因此被 Greypanel CDN 网络标记。Greypanel CDN 显示此页面以供您验证，而不是真实的人本网站的流量来源，而不是恶意软件",
                "TITLE_2": "只需点击",
                "TITLE_3": "我不是机器人",
                "TITLE_4": "用于通过安全检查的复选框。Greypanel CDN 会记住您，并且不会再次显示此页面。我们建议您在计算机上运行病毒和恶意软件扫描以消除任何感染。",
                "ROBOT": "我不是机器人",
            },
            "taiwanese": {
                "LANG": "語言",
                "CHINESE": "简体中文",
                "TAIWANESE": "繁體中文",
                "ENGLISH": "English",
                "TITLE_0": "為什麼我會看到這個頁面？",
                "TITLE_1": "我該怎麼辦",
                "PARA_0": "您正在訪問的網站受到 Greypanel CDN 的保護和加速。您的計算機可能已被惡意軟件感染，因此被 Greypanel CDN 網絡標記。 Greypanel CDN 顯示此頁面以供您驗證，而不是真實的人本網站的流量來源，而不是惡意軟件",
                "TITLE_2": "只需點擊",
                "TITLE_3": "我不是機器人",
                "TITLE_4": "用於通過安全檢查的複選框。 Greypanel CDN 會記住您，並且不會再次顯示此頁面。我們建議您在計算機上運行病毒和惡意軟件掃描以消除任何感染。",
                "ROBOT": "我不是機器人"
            },
            "english": {
                "LANG": "Language",
                "CHINESE": "简体中文",
                "TAIWANESE": "繁體中文",
                "ENGLISH": "English",
                "TITLE_0": "Why am i seeing this page?",
                "TITLE_1": "What should i do",
                "PARA_0": "The website you are visiting is protected and accelerated by Greypanel CDN. Your computer may have been infected by malware and therefore flagged by the Greypanel CDN network. Greypanel CDN displays this page for you to verify than an actual human is the source of the traffic to this site, and not malicious software.",
                "TITLE_2": "Just click the",
                "TITLE_3": "Im not a robot",
                "TITLE_4": "checkbox to pass the security check. Greypanel CDN will remember you and will not show this page again. We recommended you run a virus and malware scan on your computer to remove any infection.",
                "ROBOT": "Im not a robot"
            }
        }

        /**
         * @code 语言代码 // english, chinese
         * @langKey 对应翻译的key // zh-cn, en-gb
         */
        function convertToIden(code, langKey) {
            let id;
            if (code) {
                for (let i in identifiers) {
                    const isTrue = i === code;
                    if (isTrue) {
                        id = identifiers[i][0];
                    }
                }
            }
            if (langKey) {
                for (let i in identifiers) {
                    const isTrue = identifiers[i].some((x) => x === langKey);
                    if (isTrue) {
                        id = i;
                    }
                }
            }
            return id;
        }

        /**
         * 索取默认语言代码
         */
        function getDefinedLangKey() {
            let langKey;
            const cacheCode = utilService.add("get", 1);
            if (cacheCode) {
                langKey = cacheCode;
            }
            else if (navigator.language) {
                const flatted = Object.values(identifiers).flat();
                // other than the covered lang it might can show other language code result in not showing the related package
                if (flatted.includes(navigator.language)) {
                    return langKey = "en-us"
                }
                return langKey = navigator.language.toLowerCase()
            }
            else {
                langKey = "zh-cn"
            }
            return langKey;
        }
        // old data-sitekey：6Lf-3MwjAAAAACW20ay3db8DrwPvXDWPZLttTj7d
        function paintCaptcha(_l) {
            const content = `<div class="g-recaptcha" data-sitekey="6Lfo1WIpAAAAAMYwd5eaRUCYWtfggy6vyLt4PWwW" data-action="submit" data-callback="onSubmit"></div>`;
            const el = document.createElement("div");
            el.innerHTML = content;
            captchaEl.appendChild(el);

            // const _ee = document.getElementsByClassName("g-recaptcha")[0]

            // window.onloadCallback = function() {
            //     grecaptcha.render(_ee, {
            //         'sitekey' : '6Ld4TsgjAAAAADhIE_ANIFC9p8aZaT_UiY22hXoG',
            //         'hl': _l,
            //     });
            // }
        }


        function injectCaptcha(render) {
            const _spt = getDefinedLangKey().split("-");
            console.log(_spt);
            const _l = _spt.length >= 2 ? _spt[0]+ "-" +_spt[1].toUpperCase() : _spt[0];
            console.log(_l, 'rendered language')
            if(captchaEl) {
                paintCaptcha(_l);
                utilService.addScript(`https://www.recaptcha.net/recaptcha/api.js` )
            }

        }


        /*
         * setter getters
         */
        function contentHandler(x) {
            return element["content"].value = x.value;
        }
        function inputHandler(x, special) {
            return element["SELECT_EL"].newValue = special ? x : x.target.value;
        }
        function selectHandler(x, special) {
            const _v = special ? x : x.getAttribute("value");
            return element["selected"].value = _v;
        }
        function toggleDropdown(x) {
            return element["dropdownStatus"].newValue = x;
        }
        function getDropdownStatus() {
            return element["dropdownStatus"].value
        }
        function resizeHandler(x) {
            return element["windowWidth"].newValue = x;
        }
        function getResizeProp() {
            return element["windowWidth"].newValue
        }


        /*
         * el related
         */
        function getSelectHideEL(hide) {
            const disp = convertToIden(null, utilService.add("get", 1))
            const parsed = JSON.parse(utilService.add("get", 2));
            // console.log(parsed, 'parsed')
            const cls = hide ? `"ant-select-selection-placeholder hide"` : "ant-select-selection-placeholder";
            return `<span class=${cls}>${disp ? parsed[disp.toUpperCase()] : "Select language"}</span>`
        }


        function initBind() {
            document.querySelector("#input_s").addEventListener('input', inputHandler);

            document.querySelector("#input_s").addEventListener('click', function (event) {
                event.preventDefault();
                event.stopImmediatePropagation();
                toggleDropdown(true)
            })
            document.querySelector("body").addEventListener('click', function (event) {
                toggleDropdown(false)
            })
        }

        /**
         * @dropdownS 是否展示下拉
         * @showLangNode 是否显示language
         */
        function runEffect(dropdownS, showLangNode) {
            const langKey = getDefinedLangKey()
            // 对语言代码进行转换 索取对应的语言包
            const transformValue = convertToIden(null, langKey);

            if (langKey) {
                utilService.add('set', 1, langKey);
            }

            if (element) {
                const d = showLangNode ? `<span>${intl[transformValue]["LANG"]}</span>` : "";
                languageEl.style.display = showLangNode ? "inline-block" : "none";
                // console.log(languageEl, footerEl)
                utilService.paintToEL(languageEl, d);


                utilService.paintToEL(footerEl, `<div> Copyright © ${new Date().getFullYear()} Greypanel. All rights reserved </div>`)
                element["content"].value = transformValue;
                // set cache language package
                utilService.add("set", 2, JSON.stringify(intl[transformValue]))

                element["robot"].value = transformValue;

                element["SELECT_EL"].newValue = "";
                if (dropdownS) {
                    element["selected"].value = transformValue;
                }
            }
        }

        function initCss() {
            var style = document.createElement('style');
            style.innerHTML = `
                    .main {
                        display: flex;
                        flex-flow: column;
                        gap: 50px;
                    }
                    #order2 {
                        margin-top: 0px;
                    }
            	    #footer {
				        position: absolute;
				        bottom: 20px;
				        text-align: center;
				        left: 0;
				        right: 0;
				    }

				    select {
				        height: 40px;
				    }

				    input {
				        padding: 0;
				        background: transparent;
				        border: none;
				        outline: none;
				        -webkit-appearance: none;
				        -moz-appearance: none;
				        appearance: none;
				    }

				    span {
				        font-weight: bold;
				    }

				    body {
				        padding: 34px;
				        margin: 0;
				        color: rgba(0, 0, 0, 0.85);
				        font-size: 14px;
				        font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, 'Noto Sans', sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol', 'Noto Color Emoji';
				        font-variant: tabular-nums;
				        line-height: 1.5715;
				        background-color: #fff;
				        font-feature-settings: 'tnum';
				   
				    }

				    .txt_size {
				        font-size: 16px;
				    }

				    .options {
				        padding: 5px 10px;
				    }

				    .options:hover {
				        background-color: #f5f5f5;
				        transition: all 0.3s cubic-bezier(0.645, 0.045, 0.355, 1);
				        cursor: pointer;
				    }

				    .ant-select {
				        box-sizing: border-box;
				        margin: 0;
				        padding: 0;
				        color: rgba(0, 0, 0, 0.85);
				        font-size: 14px;
				        font-variant: tabular-nums;
				        line-height: 1.5715;
				        list-style: none;
				        font-feature-settings: 'tnum';
				        position: relative;
				        display: inline-block;
				        cursor: pointer;
				    }

				    .ant-select-selector {
				        position: relative;
				        background-color: #fff;
				        border: 1px solid #d9d9d9;
				        border-radius: 2px;
				        transition: all 0.3s cubic-bezier(0.645, 0.045, 0.355, 1);
				    }

				    .ant-select-single .ant-select-selector .ant-select-selection-search {
				        position: absolute;
				        top: 0;
				        right: 11px;
				        bottom: 0;
				        left: 11px;
				    }

				    .ant-select-single.ant-select-show-arrow .ant-select-selection-search {
				        /*right: 25px;*/
				    }

				    .ant-select:not(.ant-select-customize-input) .ant-select-selector .ant-select-selection-search-input {
				        margin: 0;
				        padding: 0;
				        background: transparent;
				        border: none;
				        outline: none;
				        -webkit-appearance: none;
				        -moz-appearance: none;
				        appearance: none;
				    }

				    .ant-select-show-search.ant-select:not(.ant-select-customize-input) .ant-select-selector input {
				        cursor: auto;
				    }

				    .ant-select-single .ant-select-selector .ant-select-selection-placeholder {
				        transition: none;
				        pointer-events: none;
				    }

				    .ant-select-single.ant-select-show-arrow .ant-select-selection-placeholder {
				        padding-right: 18px;
				    }

				    .ant-select:not(.ant-select-customize-input) .ant-select-selector {
				        position: relative;
				        background-color: #fff;
				        border: 1px solid #d9d9d9;
				        border-radius: 2px;
				        transition: all 0.3s cubic-bezier(0.645, 0.045, 0.355, 1);
				    }

				    .ant-select-show-search.ant-select:not(.ant-select-customize-input) .ant-select-selector {
				        cursor: text;
				    }

				    .ant-select-single:not(.ant-select-customize-input) .ant-select-selector {
				        width: 120px;
				        height: 32px;
				        /*padding: 0 11px;*/
				    }

				    .ant-select-single .ant-select-selector {
				        display: flex;
				    }

				    .ant-select-single:not(.ant-select-customize-input) .ant-select-selector .ant-select-selection-search-input {
				        height: 30px;
				    }

				    .ant-select-single.ant-select-show-arrow .ant-select-selection-placeholder {
				        padding-right: 18px;
				    }

				    .ant-select-single .ant-select-selector .ant-select-selection-placeholder {
				        transition: none;
				        pointer-events: none;
				    }

				    .ant-select-single .ant-select-selector .ant-select-selection-placeholder {
				        padding: 0;
				        line-height: 30px;
				        transition: all 0.3s;
				    }

				    .ant-select-selection-placeholder {
				        flex: 1;
				        overflow: hidden;
				        color: #bfbfbf;
				        white-space: nowrap;
				        text-overflow: ellipsis;
				        pointer-events: none;

				    }

				    .ant-select-arrow {
				        display: inline-block;
				        color: inherit;
				        font-style: normal;
				        line-height: 0;
				        text-transform: none;
				        vertical-align: -0.125em;
				        text-rendering: optimizelegibility;
				        -webkit-font-smoothing: antialiased;
				        -moz-osx-font-smoothing: grayscale;
				        position: absolute;
				        top: 50%;
				        right: 11px;
				        width: 12px;
				        height: 12px;
				        margin-top: -6px;
				        color: rgba(0, 0, 0, 0.25);
				        font-size: 12px;
				        line-height: 1;
				        text-align: center;
				        pointer-events: none;
				    }

				    .ant-select-single.ant-select-show-arrow .ant-select-selection-placeholder {
				        padding-right: 18px;
				    }

				    .container_data {
				        position: relative;
				    }

				    .hide {
				        visibility: hidden;
				    }


				    .decoyStyle {
				       max-width: 700px;
				        margin: auto;
				      
                        display: flex;
                        justify-content: center;
				    }

				    #select_el_hide {
				        padding-left: 10px;
				    }

				    @media only screen and (min-width: 500px) {
				        .ant-select {
				            width: 140px;
				        }

				        #container {
				            display: flex;
				            align-items: center;
				            gap: 10px;
				            justify-content: flex-end;
				        }

				        .container_data {
				            width: 140px;
				            position: absolute;
				            right: 68px;
				        }

				        .ant-select-single:not(.ant-select-customize-input) .ant-select-selector {
				            width: 100%;
				        }

				        .ant-select-dropdown {
				            width: 100%;
				            margin: 0;
				            padding: 0;
				            color: rgba(0, 0, 0, 0.85);
				            font-variant: tabular-nums;
				            line-height: 1.5715;
				            list-style: none;
				            font-feature-settings: 'tnum';
				            position: absolute;
				            z-index: 1050;
				            box-sizing: border-box;
				            padding: 4px 0;
				            overflow: hidden;
				            font-size: 14px;
				            font-variant: initial;
				            background-color: #fff;
				            border-radius: 2px;
				            outline: none;
				            box-shadow: 0 3px 6px -4px rgb(0 0 0 / 12%), 0 6px 16px 0 rgb(0 0 0 / 8%), 0 9px 28px 8px rgb(0 0 0 / 5%);
				            left: 50%;
				            transform: translateX(-50%);
				        }
				    }

				    @media only screen and (max-width: 600px) {
				        #order1 {
                            order: 3;
                        }

				        #container {
				            display: flex;
				            gap: 10px;
				            align-items: center;
				            justify-content: flex-end;
				          
				        }

				        .slt {
				            display: flex;
				            flex-direction: column-reverse;
				        }

				        .container_data {
				            width: 100%;
				            position: relative;
				            left: 50%;
				            transform: translateX(-50%);
				            top: -70px;
				        }

				        .width100 {
				            width: 100%;
				        }


				        .ant-select-single:not(.ant-select-customize-input) .ant-select-selector {
				            width: 100%;
				        }

				        .ant-select-dropdown {
				            width: 100%;
				            margin: 0;
				            padding: 0;
				            color: rgba(0, 0, 0, 0.85);
				            font-variant: tabular-nums;
				            line-height: 1.5715;
				            list-style: none;
				            font-feature-settings: 'tnum';
				            position: absolute;
				            z-index: 1050;
				            box-sizing: border-box;
				            padding: 4px 0;
				            overflow: hidden;
				            font-size: 14px;
				            font-variant: initial;
				            background-color: #fff;
				            border-radius: 2px;
				            outline: none;
				            box-shadow: 0 3px 6px -4px rgb(0 0 0 / 12%), 0 6px 16px 0 rgb(0 0 0 / 8%), 0 9px 28px 8px rgb(0 0 0 / 5%);
				            left: 50%;
				            /*top: 80px;*/
				            transform: translateX(-50%);
				            top: 40px;

				        }
				    }
                    ._btn{
                        height: 86px;
                        width: 350px;
                        border: 1px solid #D7DBE4;
                        background: #f9f9f9;
                        box-shadow: rgba(149, 157, 165, 0.2) 0px 8px 24px;
                    }

                    .checkBox{
                       /* position: absolute; */
                       display: flex;
                       margin-top: 10px;
                    }

                    .img1{
                    padding-left: 16px;

                    }
                    .text{
                        padding-left: 15px;
                        line-height: 29px;
                        font-size: 17px;
                        text-align: left;
                        color: #474747;
                    }
                    .img2{
                        padding-left: 100px;
                    }

                    ._btn:hover{
                        box-shadow: rgba(0, 0, 0, 0.1) 0px 4px 12px;
                    }
            	`
            var ref = document.querySelector('style');
            ref.parentNode.insertBefore(style, ref);
        }
        window.fnSelectHandler = selectHandler;
        window.fntoggleDropdown = toggleDropdown;
        window.onresize = function () {
            return resizeHandler(window.innerWidth)
        };
        // start
        runEffect(null, exp);
        initBind();
        initCss();
        // injectCaptcha(false);
        try {
            var noRecapDescFlagEl = document.getElementById("noRecapDescFlag");
            if(!noRecapDescFlagEl || noRecapDescFlagEl.value != '1') {
                document.getElementById('order3').style.display = ''
                document.getElementById('footer').style.display = ''
            }
        } catch(e) {
            console.error(e)
        }
    }()
}