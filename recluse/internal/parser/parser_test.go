package parser

import (
	"fmt"
	"testing"

	"github.com/scott-mescudi/eximius/recluse/pkg/models"
)


func TestParser(t *testing.T) {
	var d = &models.Document{
		Url:  "iusrhe",
		Text: document,
	}

	ParseDocument(d)

	fmt.Println(d.Text, d.Description)
}

func BenchmarkParser(b *testing.B) {
	var d = &models.Document{
		Url:  "iusrhe",
		Text: document,
	}

	for b.Loop() {
		ParseDocument(d)

	}
}

var document = `
<!DOCTYPE html>
<html class="client-nojs vector-feature-language-in-header-enabled vector-feature-language-in-main-page-header-disabled vector-feature-page-tools-pinned-disabled vector-feature-toc-pinned-clientpref-1 vector-feature-main-menu-pinned-disabled vector-feature-limited-width-clientpref-1 vector-feature-limited-width-content-enabled vector-feature-custom-font-size-clientpref-1 vector-feature-appearance-pinned-clientpref-1 vector-feature-night-mode-enabled skin-theme-clientpref-day vector-sticky-header-enabled vector-toc-available" lang="en" dir="ltr">
<head>
<meta charset="UTF-8">
<title>City, University of London - Wikipedia</title>
<script>(function(){var className="client-js vector-feature-language-in-header-enabled vector-feature-language-in-main-page-header-disabled vector-feature-page-tools-pinned-disabled vector-feature-toc-pinned-clientpref-1 vector-feature-main-menu-pinned-disabled vector-feature-limited-width-clientpref-1 vector-feature-limited-width-content-enabled vector-feature-custom-font-size-clientpref-1 vector-feature-appearance-pinned-clientpref-1 vector-feature-night-mode-enabled skin-theme-clientpref-day vector-sticky-header-enabled vector-toc-available";var cookie=document.cookie.match(/(?:^|; )enwikimwclientpreferences=([^;]+)/);if(cookie){cookie[1].split('%2C').forEach(function(pref){className=className.replace(new RegExp('(^| )'+pref.replace(/-clientpref-\w+$|[^\w-]+/g,'')+'-clientpref-\\w+( |$)'),'$1'+pref+'$2');});}document.documentElement.className=className;}());RLCONF={"wgBreakFrames":false,"wgSeparatorTransformTable":["",""],"wgDigitTransformTable":["",""],"wgDefaultDateFormat":"dmy","wgMonthNames":["","January","February","March","April","May","June","July","August","September","October","November","December"],"wgRequestId":"06a99128-2e56-415c-95a0-572d44f5ea62","wgCanonicalNamespace":"","wgCanonicalSpecialPageName":false,"wgNamespaceNumber":0,"wgPageName":"City,_University_of_London","wgTitle":"City, University of London","wgCurRevisionId":1283711199,"wgRevisionId":1283711199,"wgArticleId":533484,"wgIsArticle":true,"wgIsRedirect":false,"wgAction":"view","wgUserName":null,"wgUserGroups":["*"],"wgCategories":["Pages using gadget WikiMiniAtlas","All articles with dead external links","Articles with dead external links from November 2017","Articles with permanently dead external links","Articles with short description","Short description is different from Wikidata","Use dmy dates from December 2019","Articles to be merged from March 2025","All articles to be merged","Coordinates on Wikidata","Articles using infobox university","Pages using infobox university with the affiliations parameter","All articles with unsourced statements","Articles with unsourced statements from October 2023","Articles with unsourced statements from December 2010","Commons category link from Wikidata","Webarchive template wayback links","City, University of London","Optometry schools","Schools of informatics","Universities and colleges established in 1894","1894 establishments in England","Venues of the 1908 Summer Olympics","Olympic boxing venues","Universities UK"],"wgPageViewLanguage":"en","wgPageContentLanguage":"en","wgPageContentModel":"wikitext","wgRelevantPageName":"City,_University_of_London","wgRelevantArticleId":533484,"wgIsProbablyEditable":true,"wgRelevantPageIsProbablyEditable":true,"wgRestrictionEdit":[],"wgRestrictionMove":[],"wgNoticeProject":"wikipedia","wgCiteReferencePreviewsActive":false,"wgFlaggedRevsParams":{"tags":{"status":{"levels":1}}},"wgMediaViewerOnClick":true,"wgMediaViewerEnabledByDefault":true,"wgPopupsFlags":0,"wgVisualEditor":{"pageLanguageCode":"en","pageLanguageDir":"ltr","pageVariantFallbacks":"en"},"wgMFDisplayWikibaseDescriptions":{"search":true,"watchlist":true,"tagline":false,"nearby":true},"wgWMESchemaEditAttemptStepOversample":false,"wgWMEPageLength":50000,"wgCoordinates":{"lat":51.5278,"lon":-0.1023},"wgEditSubmitButtonLabelPublish":true,"wgULSPosition":"interlanguage","wgULSisCompactLinksEnabled":false,"wgVector2022LanguageInHeader":true,"wgULSisLanguageSelectorEmpty":false,"wgWikibaseItemId":"Q1094046","wgCheckUserClientHintsHeadersJsApi":["brands","architecture","bitness","fullVersionList","mobile","model","platform","platformVersion"],"GEHomepageSuggestedEditsEnableTopics":true,"wgGETopicsMatchModeEnabled":false,"wgGELevelingUpEnabledForUser":false};
RLSTATE={"ext.globalCssJs.user.styles":"ready","site.styles":"ready","user.styles":"ready","ext.globalCssJs.user":"ready","user":"ready","user.options":"loading","ext.cite.styles":"ready","skins.vector.search.codex.styles":"ready","skins.vector.styles":"ready","skins.vector.icons":"ready","jquery.makeCollapsible.styles":"ready","ext.wikimediamessages.styles":"ready","ext.visualEditor.desktopArticleTarget.noscript":"ready","ext.uls.interlanguage":"ready","wikibase.client.init":"ready"};RLPAGEMODULES=["ext.cite.ux-enhancements","mediawiki.page.media","site","mediawiki.page.ready","jquery.makeCollapsible","mediawiki.toc","skins.vector.js","ext.centralNotice.geoIP","ext.centralNotice.startUp","ext.gadget.ReferenceTooltips","ext.gadget.switcher","ext.gadget.WikiMiniAtlas","ext.urlShortener.toolbar","ext.centralauth.centralautologin","mmv.bootstrap","ext.popups","ext.visualEditor.desktopArticleTarget.init","ext.visualEditor.targetLoader","ext.echo.centralauth","ext.eventLogging","ext.wikimediaEvents","ext.navigationTiming","ext.uls.interface","ext.cx.eventlogging.campaigns","ext.cx.uls.quick.actions","wikibase.client.vector-2022","ext.checkUser.clientHints","ext.quicksurveys.init","ext.growthExperiments.SuggestedEditSession"];</script>
<script>(RLQ=window.RLQ||[]).push(function(){mw.loader.impl(function(){return["user.options@12s5i",function($,jQuery,require,module){mw.user.tokens.set({"patrolToken":"+\\","watchToken":"+\\","csrfToken":"+\\"});
}];});});</script>
<link rel="stylesheet" href="/w/load.php?lang=en&amp;modules=ext.cite.styles%7Cext.uls.interlanguage%7Cext.visualEditor.desktopArticleTarget.noscript%7Cext.wikimediamessages.styles%7Cjquery.makeCollapsible.styles%7Cskins.vector.icons%2Cstyles%7Cskins.vector.search.codex.styles%7Cwikibase.client.init&amp;only=styles&amp;skin=vector-2022">
<script async="" src="/w/load.php?lang=en&amp;modules=startup&amp;only=scripts&amp;raw=1&amp;skin=vector-2022"></script>
<meta name="ResourceLoaderDynamicStyles" content="">
<link rel="stylesheet" href="/w/load.php?lang=en&amp;modules=site.styles&amp;only=styles&amp;skin=vector-2022">
<meta name="generator" content="MediaWiki 1.45.0-wmf.2">
<meta name="referrer" content="origin">
<meta name="referrer" content="origin-when-cross-origin">
<meta name="robots" content="max-image-preview:standard">
<meta name="format-detection" content="telephone=no">
<meta property="og:image" content="https://upload.wikimedia.org/wikipedia/commons/thumb/0/01/Arms_of_City%2C_University_of_London.svg/1200px-Arms_of_City%2C_University_of_London.svg.png">
<meta property="og:image:width" content="1200">
<meta property="og:image:height" content="1605">
<meta property="og:image" content="https://upload.wikimedia.org/wikipedia/commons/thumb/0/01/Arms_of_City%2C_University_of_London.svg/800px-Arms_of_City%2C_University_of_London.svg.png">
<meta property="og:image:width" content="800">
<meta property="og:image:height" content="1070">
<meta property="og:image" content="https://upload.wikimedia.org/wikipedia/commons/thumb/0/01/Arms_of_City%2C_University_of_London.svg/640px-Arms_of_City%2C_University_of_London.svg.png">
<meta property="og:image:width" content="640">
<meta property="og:image:height" content="856">
<meta name="viewport" content="width=1120">
<meta property="og:title" content="City, University of London - Wikipedia">
<meta property="og:type" content="website">
<link rel="preconnect" href="//upload.wikimedia.org">
<link rel="alternate" media="only screen and (max-width: 640px)" href="//en.m.wikipedia.org/wiki/City,_University_of_London">
<link rel="alternate" type="application/x-wiki" title="Edit this page" href="/w/index.php?title=City,_University_of_London&amp;action=edit">
<link rel="apple-touch-icon" href="/static/apple-touch/wikipedia.png">
<link rel="icon" href="/static/favicon/wikipedia.ico">
<link rel="search" type="application/opensearchdescription+xml" href="/w/rest.php/v1/search" title="Wikipedia (en)">
<link rel="EditURI" type="application/rsd+xml" href="//en.wikipedia.org/w/api.php?action=rsd">
<link rel="canonical" href="https://en.wikipedia.org/wiki/City,_University_of_London">
<link rel="license" href="https://creativecommons.org/licenses/by-sa/4.0/deed.en">
<link rel="alternate" type="application/atom+xml" title="Wikipedia Atom feed" href="/w/index.php?title=Special:RecentChanges&amp;feed=atom">
<link rel="dns-prefetch" href="//meta.wikimedia.org" />
<link rel="dns-prefetch" href="auth.wikimedia.org">
</head>
<body class="skin--responsive skin-vector skin-vector-search-vue mediawiki ltr sitedir-ltr mw-hide-empty-elt ns-0 ns-subject mw-editable page-City_University_of_London rootpage-City_University_of_London skin-vector-2022 action-view"><a class="mw-jump-link" href="#bodyContent">Jump to content</a>
<div class="vector-header-container">
	<header class="vector-header mw-header no-font-mode-scale">
		<div class="vector-header-start">
			<nav class="vector-main-menu-landmark" aria-label="Site">
				
<div id="vector-main-menu-dropdown" class="vector-dropdown vector-main-menu-dropdown vector-button-flush-left vector-button-flush-right"  title="Main menu" >
	<input type="checkbox" id="vector-main-menu-dropdown-checkbox" role="button" aria-haspopup="true" data-event-name="ui.dropdown-vector-main-menu-dropdown" class="vector-dropdown-checkbox "  aria-label="Main menu"  >
	<label id="vector-main-menu-dropdown-label" for="vector-main-menu-dropdown-checkbox" class="vector-dropdown-label cdx-button cdx-button--fake-button cdx-button--fake-button--enabled cdx-button--weight-quiet cdx-button--icon-only " aria-hidden="true"  ><span class="vector-icon mw-ui-icon-menu mw-ui-icon-wikimedia-menu"></span>

<span class="vector-dropdown-label-text">Main menu</span>
	</label>
	<div class="vector-dropdown-content">


				<div id="vector-main-menu-unpinned-container" class="vector-unpinned-container">
		
<div id="vector-main-menu" class="vector-main-menu vector-pinnable-element">
	<div
	class="vector-pinnable-header vector-main-menu-pinnable-header vector-pinnable-header-unpinned"
	data-feature-name="main-menu-pinned"
	data-pinnable-element-id="vector-main-menu"
	data-pinned-container-id="vector-main-menu-pinned-container"
	data-unpinned-container-id="vector-main-menu-unpinned-container"
>
	<div class="vector-pinnable-header-label">Main menu</div>
	<button class="vector-pinnable-header-toggle-button vector-pinnable-header-pin-button" data-event-name="pinnable-header.vector-main-menu.pin">move to sidebar</button>
	<button class="vector-pinnable-header-toggle-button vector-pinnable-header-unpin-button" data-event-name="pinnable-header.vector-main-menu.unpin">hide</button>
</div>

	
<div id="p-navigation" class="vector-menu mw-portlet mw-portlet-navigation"  >
	<div class="vector-menu-heading">
		Navigation
	</div>
	<div class="vector-menu-content">
		
		<ul class="vector-menu-content-list">
			
			<li id="n-mainpage-description" class="mw-list-item"><a href="/wiki/Main_Page" title="Visit the main page [z]" accesskey="z"><span>Main page</span></a></li><li id="n-contents" class="mw-list-item"><a href="/wiki/Wikipedia:Contents" title="Guides to browsing Wikipedia"><span>Contents</span></a></li><li id="n-currentevents" class="mw-list-item"><a href="/wiki/Portal:Current_events" title="Articles related to current events"><span>Current events</span></a></li><li id="n-randompage" class="mw-list-item"><a href="/wiki/Special:Random" title="Visit a randomly selected article [x]" accesskey="x"><span>Random article</span></a></li><li id="n-aboutsite" class="mw-list-item"><a href="/wiki/Wikipedia:About" title="Learn about Wikipedia and how it works"><span>About Wikipedia</span></a></li><li id="n-contactpage" class="mw-list-item"><a href="//en.wikipedia.org/wiki/Wikipedia:Contact_us" title="How to contact Wikipedia"><span>Contact us</span></a></li>
		</ul>
		
	</div>
</div>

	
	
<div id="p-interaction" class="vector-menu mw-portlet mw-portlet-interaction"  >
	<div class="vector-menu-heading">
		Contribute
	</div>
	<div class="vector-menu-content">
		
		<ul class="vector-menu-content-list">
			
			<li id="n-help" class="mw-list-item"><a href="/wiki/Help:Contents" title="Guidance on how to use and edit Wikipedia"><span>Help</span></a></li><li id="n-introduction" class="mw-list-item"><a href="/wiki/Help:Introduction" title="Learn how to edit Wikipedia"><span>Learn to edit</span></a></li><li id="n-portal" class="mw-list-item"><a href="/wiki/Wikipedia:Community_portal" title="The hub for editors"><span>Community portal</span></a></li><li id="n-recentchanges" class="mw-list-item"><a href="/wiki/Special:RecentChanges" title="A list of recent changes to Wikipedia [r]" accesskey="r"><span>Recent changes</span></a></li><li id="n-upload" class="mw-list-item"><a href="/wiki/Wikipedia:File_upload_wizard" title="Add images or other media for use on Wikipedia"><span>Upload file</span></a></li><li id="n-specialpages" class="mw-list-item"><a href="/wiki/Special:SpecialPages"><span>Special pages</span></a></li>
		</ul>
		
	</div>
</div>

</div>

				</div>

	</div>
</div>

		</nav>
			
<a href="/wiki/Main_Page" class="mw-logo">
	<img class="mw-logo-icon" src="/static/images/icons/wikipedia.png" alt="" aria-hidden="true" height="50" width="50">
	<span class="mw-logo-container skin-invert">
		<img class="mw-logo-wordmark" alt="Wikipedia" src="/static/images/mobile/copyright/wikipedia-wordmark-en.svg" style="width: 7.5em; height: 1.125em;">
		<img class="mw-logo-tagline" alt="The Free Encyclopedia" src="/static/images/mobile/copyright/wikipedia-tagline-en.svg" width="117" height="13" style="width: 7.3125em; height: 0.8125em;">
	</span>
</a>

		</div>
		<div class="vector-header-end">
			
<div id="p-search" role="search" class="vector-search-box-vue  vector-search-box-collapses vector-search-box-show-thumbnail vector-search-box-auto-expand-width vector-search-box">
	<a href="/wiki/Special:Search" class="cdx-button cdx-button--fake-button cdx-button--fake-button--enabled cdx-button--weight-quiet cdx-button--icon-only search-toggle" title="Search Wikipedia [f]" accesskey="f"><span class="vector-icon mw-ui-icon-search mw-ui-icon-wikimedia-search"></span>

<span>Search</span>
	</a>
	<div class="vector-typeahead-search-container">
		<div class="cdx-typeahead-search cdx-typeahead-search--show-thumbnail cdx-typeahead-search--auto-expand-width">
			<form action="/w/index.php" id="searchform" class="cdx-search-input cdx-search-input--has-end-button">
				<div id="simpleSearch" class="cdx-search-input__input-wrapper"  data-search-loc="header-moved">
					<div class="cdx-text-input cdx-text-input--has-start-icon">
						<input
							class="cdx-text-input__input mw-searchInput"
							 type="search" name="search" placeholder="Search Wikipedia" aria-label="Search Wikipedia" autocapitalize="sentences" spellcheck="false" title="Search Wikipedia [f]" accesskey="f" id="searchInput"
							>
						<span class="cdx-text-input__icon cdx-text-input__start-icon"></span>
					</div>
					<input type="hidden" name="title" value="Special:Search">
				</div>
				<button class="cdx-button cdx-search-input__end-button">Search</button>
			</form>
		</div>
	</div>
</div>

			<nav class="vector-user-links vector-user-links-wide" aria-label="Personal tools">
	<div class="vector-user-links-main">
	
<div id="p-vector-user-menu-preferences" class="vector-menu mw-portlet emptyPortlet"  >
	<div class="vector-menu-content">
		
		<ul class="vector-menu-content-list">
			
			
		</ul>
		
	</div>
</div>

	
<div id="p-vector-user-menu-userpage" class="vector-menu mw-portlet emptyPortlet"  >
	<div class="vector-menu-content">
		
		<ul class="vector-menu-content-list">
			
			
		</ul>
		
	</div>
</div>

	<nav class="vector-appearance-landmark" aria-label="Appearance">
		
<div id="vector-appearance-dropdown" class="vector-dropdown "  title="Change the appearance of the page&#039;s font size, width, and color" >
	<input type="checkbox" id="vector-appearance-dropdown-checkbox" role="button" aria-haspopup="true" data-event-name="ui.dropdown-vector-appearance-dropdown" class="vector-dropdown-checkbox "  aria-label="Appearance"  >
	<label id="vector-appearance-dropdown-label" for="vector-appearance-dropdown-checkbox" class="vector-dropdown-label cdx-button cdx-button--fake-button cdx-button--fake-button--enabled cdx-button--weight-quiet cdx-button--icon-only " aria-hidden="true"  ><span class="vector-icon mw-ui-icon-appearance mw-ui-icon-wikimedia-appearance"></span>

<span class="vector-dropdown-label-text">Appearance</span>
	</label>
	<div class="vector-dropdown-content">


			<div id="vector-appearance-unpinned-container" class="vector-unpinned-container">
				
			</div>
		
	</div>
</div>

	</nav>
	
<div id="p-vector-user-menu-notifications" class="vector-menu mw-portlet emptyPortlet"  >
	<div class="vector-menu-content">
		
		<ul class="vector-menu-content-list">
			
			
		</ul>
		
	</div>
</div>

	
<div id="p-vector-user-menu-overflow" class="vector-menu mw-portlet"  >
	<div class="vector-menu-content">
		
		<ul class="vector-menu-content-list">
			<li id="pt-sitesupport-2" class="user-links-collapsible-item mw-list-item user-links-collapsible-item"><a data-mw="interface" href="https://donate.wikimedia.org/?wmf_source=donate&amp;wmf_medium=sidebar&amp;wmf_campaign=en.wikipedia.org&amp;uselang=en" class=""><span>Donate</span></a>
</li>
<li id="pt-createaccount-2" class="user-links-collapsible-item mw-list-item user-links-collapsible-item"><a data-mw="interface" href="/w/index.php?title=Special:CreateAccount&amp;returnto=City%2C+University+of+London" title="You are encouraged to create an account and log in; however, it is not mandatory" class=""><span>Create account</span></a>
</li>
<li id="pt-login-2" class="user-links-collapsible-item mw-list-item user-links-collapsible-item"><a data-mw="interface" href="/w/index.php?title=Special:UserLogin&amp;returnto=City%2C+University+of+London" title="You&#039;re encouraged to log in; however, it&#039;s not mandatory. [o]" accesskey="o" class=""><span>Log in</span></a>
</li>

			
		</ul>
		
	</div>
</div>

	</div>
	
<div id="vector-user-links-dropdown" class="vector-dropdown vector-user-menu vector-button-flush-right vector-user-menu-logged-out"  title="Log in and more options" >
	<input type="checkbox" id="vector-user-links-dropdown-checkbox" role="button" aria-haspopup="true" data-event-name="ui.dropdown-vector-user-links-dropdown" class="vector-dropdown-checkbox "  aria-label="Personal tools"  >
	<label id="vector-user-links-dropdown-label" for="vector-user-links-dropdown-checkbox" class="vector-dropdown-label cdx-button cdx-button--fake-button cdx-button--fake-button--enabled cdx-button--weight-quiet cdx-button--icon-only " aria-hidden="true"  ><span class="vector-icon mw-ui-icon-ellipsis mw-ui-icon-wikimedia-ellipsis"></span>

<span class="vector-dropdown-label-text">Personal tools</span>
	</label>
	<div class="vector-dropdown-content">


		
<div id="p-personal" class="vector-menu mw-portlet mw-portlet-personal user-links-collapsible-item"  title="User menu" >
	<div class="vector-menu-content">
		
		<ul class="vector-menu-content-list">
			
			<li id="pt-sitesupport" class="user-links-collapsible-item mw-list-item"><a href="https://donate.wikimedia.org/?wmf_source=donate&amp;wmf_medium=sidebar&amp;wmf_campaign=en.wikipedia.org&amp;uselang=en"><span>Donate</span></a></li><li id="pt-createaccount" class="user-links-collapsible-item mw-list-item"><a href="/w/index.php?title=Special:CreateAccount&amp;returnto=City%2C+University+of+London" title="You are encouraged to create an account and log in; however, it is not mandatory"><span class="vector-icon mw-ui-icon-userAdd mw-ui-icon-wikimedia-userAdd"></span> <span>Create account</span></a></li><li id="pt-login" class="user-links-collapsible-item mw-list-item"><a href="/w/index.php?title=Special:UserLogin&amp;returnto=City%2C+University+of+London" title="You&#039;re encouraged to log in; however, it&#039;s not mandatory. [o]" accesskey="o"><span class="vector-icon mw-ui-icon-logIn mw-ui-icon-wikimedia-logIn"></span> <span>Log in</span></a></li>
		</ul>
		
	</div>
</div>

<div id="p-user-menu-anon-editor" class="vector-menu mw-portlet mw-portlet-user-menu-anon-editor"  >
	<div class="vector-menu-heading">
		Pages for logged out editors <a href="/wiki/Help:Introduction" aria-label="Learn more about editing"><span>learn more</span></a>
	</div>
	<div class="vector-menu-content">
		
		<ul class="vector-menu-content-list">
			
			<li id="pt-anoncontribs" class="mw-list-item"><a href="/wiki/Special:MyContributions" title="A list of edits made from this IP address [y]" accesskey="y"><span>Contributions</span></a></li><li id="pt-anontalk" class="mw-list-item"><a href="/wiki/Special:MyTalk" title="Discussion about edits from this IP address [n]" accesskey="n"><span>Talk</span></a></li>
		</ul>
		
	</div>
</div>

	
	</div>
</div>

</nav>

		</div>
	</header>
</div>
<div class="mw-page-container">
	<div class="mw-page-container-inner">
		<div class="vector-sitenotice-container">
			<div id="siteNotice"><!-- CentralNotice --></div>
		</div>
		<div class="vector-column-start">
			<div class="vector-main-menu-container">
		<div id="mw-navigation">
			<nav id="mw-panel" class="vector-main-menu-landmark" aria-label="Site">
				<div id="vector-main-menu-pinned-container" class="vector-pinned-container">
				
				</div>
		</nav>
		</div>
	</div>
	<div class="vector-sticky-pinned-container">
				<nav id="mw-panel-toc" aria-label="Contents" data-event-name="ui.sidebar-toc" class="mw-table-of-contents-container vector-toc-landmark">
					<div id="vector-toc-pinned-container" class="vector-pinned-container">
					<div id="vector-toc" class="vector-toc vector-pinnable-element">
	<div
	class="vector-pinnable-header vector-toc-pinnable-header vector-pinnable-header-pinned"
	data-feature-name="toc-pinned"
	data-pinnable-element-id="vector-toc"
	
	
>
	<h2 class="vector-pinnable-header-label">Contents</h2>
	<button class="vector-pinnable-header-toggle-button vector-pinnable-header-pin-button" data-event-name="pinnable-header.vector-toc.pin">move to sidebar</button>
	<button class="vector-pinnable-header-toggle-button vector-pinnable-header-unpin-button" data-event-name="pinnable-header.vector-toc.unpin">hide</button>
</div>


	<ul class="vector-toc-contents" id="mw-panel-toc-list">
		<li id="toc-mw-content-text"
			class="vector-toc-list-item vector-toc-level-1">
			<a href="#" class="vector-toc-link">
				<div class="vector-toc-text">(Top)</div>
			</a>
		</li>
		<li id="toc-History"
		class="vector-toc-list-item vector-toc-level-1">
		<a class="vector-toc-link" href="#History">
			<div class="vector-toc-text">
				<span class="vector-toc-numb">1</span>
				<span>History</span>
			</div>
		</a>
		
			<button aria-controls="toc-History-sublist" class="cdx-button cdx-button--weight-quiet cdx-button--icon-only vector-toc-toggle">
				<span class="vector-icon mw-ui-icon-wikimedia-expand"></span>
				<span>Toggle History subsection</span>
			</button>
		
		<ul id="toc-History-sublist" class="vector-toc-list">
			<li id="toc-Origins"
			class="vector-toc-list-item vector-toc-level-2">
			<a class="vector-toc-link" href="#Origins">
				<div class="vector-toc-text">
					<span class="vector-toc-numb">1.1</span>
					<span>Origins</span>
				</div>
			</a>
			
			<ul id="toc-Origins-sublist" class="vector-toc-list">
			</ul>
		</li>
		<li id="toc-20th_century"
			class="vector-toc-list-item vector-toc-level-2">
			<a class="vector-toc-link" href="#20th_century">
				<div class="vector-toc-text">
					<span class="vector-toc-numb">1.2</span>
					<span>20th century</span>
				</div>
			</a>
			
			<ul id="toc-20th_century-sublist" class="vector-toc-list">
			</ul>
		</li>
		<li id="toc-21st_century"
			class="vector-toc-list-item vector-toc-level-2">
			<a class="vector-toc-link" href="#21st_century">
				<div class="vector-toc-text">
					<span class="vector-toc-numb">1.3</span>
					<span>21st century</span>
				</div>
			</a>
			
			<ul id="toc-21st_century-sublist" class="vector-toc-list">
			</ul>
		</li>
	</ul>
	</li>
	<li id="toc-Campus"
		class="vector-toc-list-item vector-toc-level-1">
		<a class="vector-toc-link" href="#Campus">
			<div class="vector-toc-text">
				<span class="vector-toc-numb">2</span>
				<span>Campus</span>
			</div>
		</a>
		
		<ul id="toc-Campus-sublist" class="vector-toc-list">
		</ul>
	</li>
	<li id="toc-Organisation_and_administration"
		class="vector-toc-list-item vector-toc-level-1">
		<a class="vector-toc-link" href="#Organisation_and_administration">
			<div class="vector-toc-text">
				<span class="vector-toc-numb">3</span>
				<span>Organisation and administration</span>
			</div>
		</a>
		
			<button aria-controls="toc-Organisation_and_administration-sublist" class="cdx-button cdx-button--weight-quiet cdx-button--icon-only vector-toc-toggle">
				<span class="vector-icon mw-ui-icon-wikimedia-expand"></span>
				<span>Toggle Organisation and administration subsection</span>
			</button>
		
		<ul id="toc-Organisation_and_administration-sublist" class="vector-toc-list">
			<li id="toc-Schools"
			class="vector-toc-list-item vector-toc-level-2">
			<a class="vector-toc-link" href="#Schools">
				<div class="vector-toc-text">
					<span class="vector-toc-numb">3.1</span>
					<span>Schools</span>
				</div>
			</a>
			
			<ul id="toc-Schools-sublist" class="vector-toc-list">
			</ul>
		</li>
		<li id="toc-Finances"
			class="vector-toc-list-item vector-toc-level-2">
			<a class="vector-toc-link" href="#Finances">
				<div class="vector-toc-text">
					<span class="vector-toc-numb">3.2</span>
					<span>Finances</span>
				</div>
			</a>
			
			<ul id="toc-Finances-sublist" class="vector-toc-list">
			</ul>
		</li>
	</ul>
	</li>
	<li id="toc-Academic_profile"
		class="vector-toc-list-item vector-toc-level-1">
		<a class="vector-toc-link" href="#Academic_profile">
			<div class="vector-toc-text">
				<span class="vector-toc-numb">4</span>
				<span>Academic profile</span>
			</div>
		</a>
		
			<button aria-controls="toc-Academic_profile-sublist" class="cdx-button cdx-button--weight-quiet cdx-button--icon-only vector-toc-toggle">
				<span class="vector-icon mw-ui-icon-wikimedia-expand"></span>
				<span>Toggle Academic profile subsection</span>
			</button>
		
		<ul id="toc-Academic_profile-sublist" class="vector-toc-list">
			<li id="toc-Courses_and_rankings"
			class="vector-toc-list-item vector-toc-level-2">
			<a class="vector-toc-link" href="#Courses_and_rankings">
				<div class="vector-toc-text">
					<span class="vector-toc-numb">4.1</span>
					<span>Courses and rankings</span>
				</div>
			</a>
			
			<ul id="toc-Courses_and_rankings-sublist" class="vector-toc-list">
			</ul>
		</li>
		<li id="toc-Partnerships_and_collaborations"
			class="vector-toc-list-item vector-toc-level-2">
			<a class="vector-toc-link" href="#Partnerships_and_collaborations">
				<div class="vector-toc-text">
					<span class="vector-toc-numb">4.2</span>
					<span>Partnerships and collaborations</span>
				</div>
			</a>
			
			<ul id="toc-Partnerships_and_collaborations-sublist" class="vector-toc-list">
				<li id="toc-CETL"
			class="vector-toc-list-item vector-toc-level-3">
			<a class="vector-toc-link" href="#CETL">
				<div class="vector-toc-text">
					<span class="vector-toc-numb">4.2.1</span>
					<span>CETL</span>
				</div>
			</a>
			
			<ul id="toc-CETL-sublist" class="vector-toc-list">
			</ul>
		</li>
		<li id="toc-City_of_London"
			class="vector-toc-list-item vector-toc-level-3">
			<a class="vector-toc-link" href="#City_of_London">
				<div class="vector-toc-text">
					<span class="vector-toc-numb">4.2.2</span>
					<span>City of London</span>
				</div>
			</a>
			
			<ul id="toc-City_of_London-sublist" class="vector-toc-list">
			</ul>
		</li>
		<li id="toc-LCACE"
			class="vector-toc-list-item vector-toc-level-3">
			<a class="vector-toc-link" href="#LCACE">
				<div class="vector-toc-text">
					<span class="vector-toc-numb">4.2.3</span>
					<span>LCACE</span>
				</div>
			</a>
			
			<ul id="toc-LCACE-sublist" class="vector-toc-list">
			</ul>
		</li>
		<li id="toc-WC2_University_Network"
			class="vector-toc-list-item vector-toc-level-3">
			<a class="vector-toc-link" href="#WC2_University_Network">
				<div class="vector-toc-text">
					<span class="vector-toc-numb">4.2.4</span>
					<span>WC2 University Network</span>
				</div>
			</a>
			
			<ul id="toc-WC2_University_Network-sublist" class="vector-toc-list">
			</ul>
		</li>
		<li id="toc-Erasmus_Mundus_MULTI"
			class="vector-toc-list-item vector-toc-level-3">
			<a class="vector-toc-link" href="#Erasmus_Mundus_MULTI">
				<div class="vector-toc-text">
					<span class="vector-toc-numb">4.2.5</span>
					<span>Erasmus Mundus MULTI</span>
				</div>
			</a>
			
			<ul id="toc-Erasmus_Mundus_MULTI-sublist" class="vector-toc-list">
			</ul>
		</li>
		<li id="toc-UCL_Partners"
			class="vector-toc-list-item vector-toc-level-3">
			<a class="vector-toc-link" href="#UCL_Partners">
				<div class="vector-toc-text">
					<span class="vector-toc-numb">4.2.6</span>
					<span>UCL Partners</span>
				</div>
			</a>
			
			<ul id="toc-UCL_Partners-sublist" class="vector-toc-list">
			</ul>
		</li>
		<li id="toc-City_Research_Online"
			class="vector-toc-list-item vector-toc-level-3">
			<a class="vector-toc-link" href="#City_Research_Online">
				<div class="vector-toc-text">
					<span class="vector-toc-numb">4.2.7</span>
					<span>City Research Online</span>
				</div>
			</a>
			
			<ul id="toc-City_Research_Online-sublist" class="vector-toc-list">
			</ul>
		</li>
	</ul>
		</li>
	</ul>
	</li>
	<li id="toc-Student_life"
		class="vector-toc-list-item vector-toc-level-1">
		<a class="vector-toc-link" href="#Student_life">
			<div class="vector-toc-text">
				<span class="vector-toc-numb">5</span>
				<span>Student life</span>
			</div>
		</a>
		
			<button aria-controls="toc-Student_life-sublist" class="cdx-button cdx-button--weight-quiet cdx-button--icon-only vector-toc-toggle">
				<span class="vector-icon mw-ui-icon-wikimedia-expand"></span>
				<span>Toggle Student life subsection</span>
			</button>
		
		<ul id="toc-Student_life-sublist" class="vector-toc-list">
			<li id="toc-Students&#039;_Union"
			class="vector-toc-list-item vector-toc-level-2">
			<a class="vector-toc-link" href="#Students&#039;_Union">
				<div class="vector-toc-text">
					<span class="vector-toc-numb">5.1</span>
					<span>Students' Union</span>
				</div>
			</a>
			
			<ul id="toc-Students&#039;_Union-sublist" class="vector-toc-list">
			</ul>
		</li>
		<li id="toc-Student_media"
			class="vector-toc-list-item vector-toc-level-2">
			<a class="vector-toc-link" href="#Student_media">
				<div class="vector-toc-text">
					<span class="vector-toc-numb">5.2</span>
					<span>Student media</span>
				</div>
			</a>
			
			<ul id="toc-Student_media-sublist" class="vector-toc-list">
			</ul>
		</li>
		<li id="toc-Other"
			class="vector-toc-list-item vector-toc-level-2">
			<a class="vector-toc-link" href="#Other">
				<div class="vector-toc-text">
					<span class="vector-toc-numb">5.3</span>
					<span>Other</span>
				</div>
			</a>
			
			<ul id="toc-Other-sublist" class="vector-toc-list">
			</ul>
		</li>
	</ul>
	</li>
	<li id="toc-Sustainability_ranking"
		class="vector-toc-list-item vector-toc-level-1">
		<a class="vector-toc-link" href="#Sustainability_ranking">
			<div class="vector-toc-text">
				<span class="vector-toc-numb">6</span>
				<span>Sustainability ranking</span>
			</div>
		</a>
		
		<ul id="toc-Sustainability_ranking-sublist" class="vector-toc-list">
		</ul>
	</li>
	<li id="toc-Notable_people"
		class="vector-toc-list-item vector-toc-level-1">
		<a class="vector-toc-link" href="#Notable_people">
			<div class="vector-toc-text">
				<span class="vector-toc-numb">7</span>
				<span>Notable people</span>
			</div>
		</a>
		
			<button aria-controls="toc-Notable_people-sublist" class="cdx-button cdx-button--weight-quiet cdx-button--icon-only vector-toc-toggle">
				<span class="vector-icon mw-ui-icon-wikimedia-expand"></span>
				<span>Toggle Notable people subsection</span>
			</button>
		
		<ul id="toc-Notable_people-sublist" class="vector-toc-list">
			<li id="toc-Notable_alumni"
			class="vector-toc-list-item vector-toc-level-2">
			<a class="vector-toc-link" href="#Notable_alumni">
				<div class="vector-toc-text">
					<span class="vector-toc-numb">7.1</span>
					<span>Notable alumni</span>
				</div>
			</a>
			
			<ul id="toc-Notable_alumni-sublist" class="vector-toc-list">
				<li id="toc-Government,_politics_and_society"
			class="vector-toc-list-item vector-toc-level-3">
			<a class="vector-toc-link" href="#Government,_politics_and_society">
				<div class="vector-toc-text">
					<span class="vector-toc-numb">7.1.1</span>
					<span>Government, politics and society</span>
				</div>
			</a>
			
			<ul id="toc-Government,_politics_and_society-sublist" class="vector-toc-list">
			</ul>
		</li>
		<li id="toc-Arts,_science_and_academia"
			class="vector-toc-list-item vector-toc-level-3">
			<a class="vector-toc-link" href="#Arts,_science_and_academia">
				<div class="vector-toc-text">
					<span class="vector-toc-numb">7.1.2</span>
					<span>Arts, science and academia</span>
				</div>
			</a>
			
			<ul id="toc-Arts,_science_and_academia-sublist" class="vector-toc-list">
			</ul>
		</li>
		<li id="toc-Business_and_finance"
			class="vector-toc-list-item vector-toc-level-3">
			<a class="vector-toc-link" href="#Business_and_finance">
				<div class="vector-toc-text">
					<span class="vector-toc-numb">7.1.3</span>
					<span>Business and finance</span>
				</div>
			</a>
			
			<ul id="toc-Business_and_finance-sublist" class="vector-toc-list">
			</ul>
		</li>
		<li id="toc-Media_and_entertainment"
			class="vector-toc-list-item vector-toc-level-3">
			<a class="vector-toc-link" href="#Media_and_entertainment">
				<div class="vector-toc-text">
					<span class="vector-toc-numb">7.1.4</span>
					<span>Media and entertainment</span>
				</div>
			</a>
			
			<ul id="toc-Media_and_entertainment-sublist" class="vector-toc-list">
			</ul>
		</li>
	</ul>
		</li>
		<li id="toc-Notable_faculty_and_staff"
			class="vector-toc-list-item vector-toc-level-2">
			<a class="vector-toc-link" href="#Notable_faculty_and_staff">
				<div class="vector-toc-text">
					<span class="vector-toc-numb">7.2</span>
					<span>Notable faculty and staff</span>
				</div>
			</a>
			
			<ul id="toc-Notable_faculty_and_staff-sublist" class="vector-toc-list">
			</ul>
		</li>
		<li id="toc-Vice-Chancellors_(Pre-2016)_/_Presidents_(Post-2016)"
			class="vector-toc-list-item vector-toc-level-2">
			<a class="vector-toc-link" href="#Vice-Chancellors_(Pre-2016)_/_Presidents_(Post-2016)">
				<div class="vector-toc-text">
					<span class="vector-toc-numb">7.3</span>
					<span>Vice-Chancellors (Pre-2016) / Presidents (Post-2016)</span>
				</div>
			</a>
			
			<ul id="toc-Vice-Chancellors_(Pre-2016)_/_Presidents_(Post-2016)-sublist" class="vector-toc-list">
			</ul>
		</li>
	</ul>
	</li>
	<li id="toc-In_popular_culture"
		class="vector-toc-list-item vector-toc-level-1">
		<a class="vector-toc-link" href="#In_popular_culture">
			<div class="vector-toc-text">
				<span class="vector-toc-numb">8</span>
				<span>In popular culture</span>
			</div>
		</a>
		
		<ul id="toc-In_popular_culture-sublist" class="vector-toc-list">
		</ul>
	</li>
	<li id="toc-See_also"
		class="vector-toc-list-item vector-toc-level-1">
		<a class="vector-toc-link" href="#See_also">
			<div class="vector-toc-text">
				<span class="vector-toc-numb">9</span>
				<span>See also</span>
			</div>
		</a>
		
		<ul id="toc-See_also-sublist" class="vector-toc-list">
		</ul>
	</li>
	<li id="toc-References"
		class="vector-toc-list-item vector-toc-level-1">
		<a class="vector-toc-link" href="#References">
			<div class="vector-toc-text">
				<span class="vector-toc-numb">10</span>
				<span>References</span>
			</div>
		</a>
		
		<ul id="toc-References-sublist" class="vector-toc-list">
		</ul>
	</li>
	<li id="toc-External_links"
		class="vector-toc-list-item vector-toc-level-1">
		<a class="vector-toc-link" href="#External_links">
			<div class="vector-toc-text">
				<span class="vector-toc-numb">11</span>
				<span>External links</span>
			</div>
		</a>
		
		<ul id="toc-External_links-sublist" class="vector-toc-list">
		</ul>
	</li>
</ul>
</div>

					</div>
		</nav>
			</div>
		</div>
		<div class="mw-content-container">
			<main id="content" class="mw-body">
				<header class="mw-body-header vector-page-titlebar no-font-mode-scale">
					<nav aria-label="Contents" class="vector-toc-landmark">
						
<div id="vector-page-titlebar-toc" class="vector-dropdown vector-page-titlebar-toc vector-button-flush-left"  title="Table of Contents" >
	<input type="checkbox" id="vector-page-titlebar-toc-checkbox" role="button" aria-haspopup="true" data-event-name="ui.dropdown-vector-page-titlebar-toc" class="vector-dropdown-checkbox "  aria-label="Toggle the table of contents"  >
	<label id="vector-page-titlebar-toc-label" for="vector-page-titlebar-toc-checkbox" class="vector-dropdown-label cdx-button cdx-button--fake-button cdx-button--fake-button--enabled cdx-button--weight-quiet cdx-button--icon-only " aria-hidden="true"  ><span class="vector-icon mw-ui-icon-listBullet mw-ui-icon-wikimedia-listBullet"></span>

<span class="vector-dropdown-label-text">Toggle the table of contents</span>
	</label>
	<div class="vector-dropdown-content">


							<div id="vector-page-titlebar-toc-unpinned-container" class="vector-unpinned-container">
			</div>
		
	</div>
</div>

					</nav>
					<h1 id="firstHeading" class="firstHeading mw-first-heading"><span class="mw-page-title-main">City, University of London</span></h1>
							
<div id="p-lang-btn" class="vector-dropdown mw-portlet mw-portlet-lang"  >
	<input type="checkbox" id="p-lang-btn-checkbox" role="button" aria-haspopup="true" data-event-name="ui.dropdown-p-lang-btn" class="vector-dropdown-checkbox mw-interlanguage-selector" aria-label="Go to an article in another language. Available in 30 languages"   >
	<label id="p-lang-btn-label" for="p-lang-btn-checkbox" class="vector-dropdown-label cdx-button cdx-button--fake-button cdx-button--fake-button--enabled cdx-button--weight-quiet cdx-button--action-progressive mw-portlet-lang-heading-30" aria-hidden="true"  ><span class="vector-icon mw-ui-icon-language-progressive mw-ui-icon-wikimedia-language-progressive"></span>

<span class="vector-dropdown-label-text">30 languages</span>
	</label>
	<div class="vector-dropdown-content">

		<div class="vector-menu-content">
			
			<ul class="vector-menu-content-list">
				
				<li class="interlanguage-link interwiki-ar mw-list-item"><a href="https://ar.wikipedia.org/wiki/%D8%AC%D8%A7%D9%85%D8%B9%D8%A9_%D8%B3%D9%8A%D8%AA%D9%8A_%D9%84%D9%86%D8%AF%D9%86" title="جامعة سيتي لندن – Arabic" lang="ar" hreflang="ar" data-title="جامعة سيتي لندن" data-language-autonym="العربية" data-language-local-name="Arabic" class="interlanguage-link-target"><span>العربية</span></a></li><li class="interlanguage-link interwiki-az mw-list-item"><a href="https://az.wikipedia.org/wiki/London_%C5%9E%C9%99h%C9%99r_Universiteti" title="London Şəhər Universiteti – Azerbaijani" lang="az" hreflang="az" data-title="London Şəhər Universiteti" data-language-autonym="Azərbaycanca" data-language-local-name="Azerbaijani" class="interlanguage-link-target"><span>Azərbaycanca</span></a></li><li class="interlanguage-link interwiki-be mw-list-item"><a href="https://be.wikipedia.org/wiki/%D0%9B%D0%BE%D0%BD%D0%B4%D0%B0%D0%BD%D1%81%D0%BA%D1%96_%D0%B3%D0%B0%D1%80%D0%B0%D0%B4%D1%81%D0%BA%D1%96_%D1%9E%D0%BD%D1%96%D0%B2%D0%B5%D1%80%D1%81%D1%96%D1%82%D1%8D%D1%82" title="Лонданскі гарадскі ўніверсітэт – Belarusian" lang="be" hreflang="be" data-title="Лонданскі гарадскі ўніверсітэт" data-language-autonym="Беларуская" data-language-local-name="Belarusian" class="interlanguage-link-target"><span>Беларуская</span></a></li><li class="interlanguage-link interwiki-da mw-list-item"><a href="https://da.wikipedia.org/wiki/City_University_(London)" title="City University (London) – Danish" lang="da" hreflang="da" data-title="City University (London)" data-language-autonym="Dansk" data-language-local-name="Danish" class="interlanguage-link-target"><span>Dansk</span></a></li><li class="interlanguage-link interwiki-de mw-list-item"><a href="https://de.wikipedia.org/wiki/City,_University_of_London" title="City, University of London – German" lang="de" hreflang="de" data-title="City, University of London" data-language-autonym="Deutsch" data-language-local-name="German" class="interlanguage-link-target"><span>Deutsch</span></a></li><li class="interlanguage-link interwiki-et mw-list-item"><a href="https://et.wikipedia.org/wiki/Londoni_Linna%C3%BClikool" title="Londoni Linnaülikool – Estonian" lang="et" hreflang="et" data-title="Londoni Linnaülikool" data-language-autonym="Eesti" data-language-local-name="Estonian" class="interlanguage-link-target"><span>Eesti</span></a></li><li class="interlanguage-link interwiki-el mw-list-item"><a href="https://el.wikipedia.org/wiki/%CE%A0%CE%B1%CE%BD%CE%B5%CF%80%CE%B9%CF%83%CF%84%CE%AE%CE%BC%CE%B9%CE%BF_%CE%A3%CE%AF%CF%84%CE%B9_%CF%84%CE%BF%CF%85_%CE%9B%CE%BF%CE%BD%CE%B4%CE%AF%CE%BD%CE%BF%CF%85" title="Πανεπιστήμιο Σίτι του Λονδίνου – Greek" lang="el" hreflang="el" data-title="Πανεπιστήμιο Σίτι του Λονδίνου" data-language-autonym="Ελληνικά" data-language-local-name="Greek" class="interlanguage-link-target"><span>Ελληνικά</span></a></li><li class="interlanguage-link interwiki-es mw-list-item"><a href="https://es.wikipedia.org/wiki/Universidad_de_la_City_de_Londres" title="Universidad de la City de Londres – Spanish" lang="es" hreflang="es" data-title="Universidad de la City de Londres" data-language-autonym="Español" data-language-local-name="Spanish" class="interlanguage-link-target"><span>Español</span></a></li><li class="interlanguage-link interwiki-fa mw-list-item"><a href="https://fa.wikipedia.org/wiki/%D8%B3%DB%8C%D8%AA%DB%8C%D8%8C_%D8%AF%D8%A7%D9%86%D8%B4%DA%AF%D8%A7%D9%87_%D9%84%D9%86%D8%AF%D9%86" title="سیتی، دانشگاه لندن – Persian" lang="fa" hreflang="fa" data-title="سیتی، دانشگاه لندن" data-language-autonym="فارسی" data-language-local-name="Persian" class="interlanguage-link-target"><span>فارسی</span></a></li><li class="interlanguage-link interwiki-fr mw-list-item"><a href="https://fr.wikipedia.org/wiki/City_University" title="City University – French" lang="fr" hreflang="fr" data-title="City University" data-language-autonym="Français" data-language-local-name="French" class="interlanguage-link-target"><span>Français</span></a></li><li class="interlanguage-link interwiki-ko mw-list-item"><a href="https://ko.wikipedia.org/wiki/%EB%9F%B0%EB%8D%98_%EB%8C%80%ED%95%99%EA%B5%90_%EC%8B%9C%ED%8B%B0" title="런던 대학교 시티 – Korean" lang="ko" hreflang="ko" data-title="런던 대학교 시티" data-language-autonym="한국어" data-language-local-name="Korean" class="interlanguage-link-target"><span>한국어</span></a></li><li class="interlanguage-link interwiki-hy mw-list-item"><a href="https://hy.wikipedia.org/wiki/%D4%BC%D5%B8%D5%B6%D5%A4%D5%B8%D5%B6%D5%AB_%D6%84%D5%A1%D5%B2%D5%A1%D6%84%D5%A1%D5%B5%D5%AB%D5%B6_%D5%B0%D5%A1%D5%B4%D5%A1%D5%AC%D5%BD%D5%A1%D6%80%D5%A1%D5%B6" title="Լոնդոնի քաղաքային համալսարան – Armenian" lang="hy" hreflang="hy" data-title="Լոնդոնի քաղաքային համալսարան" data-language-autonym="Հայերեն" data-language-local-name="Armenian" class="interlanguage-link-target"><span>Հայերեն</span></a></li><li class="interlanguage-link interwiki-it mw-list-item"><a href="https://it.wikipedia.org/wiki/City_University_(Londra)" title="City University (Londra) – Italian" lang="it" hreflang="it" data-title="City University (Londra)" data-language-autonym="Italiano" data-language-local-name="Italian" class="interlanguage-link-target"><span>Italiano</span></a></li><li class="interlanguage-link interwiki-lv mw-list-item"><a href="https://lv.wikipedia.org/wiki/Londonas_Pils%C4%93tas_universit%C4%81te" title="Londonas Pilsētas universitāte – Latvian" lang="lv" hreflang="lv" data-title="Londonas Pilsētas universitāte" data-language-autonym="Latviešu" data-language-local-name="Latvian" class="interlanguage-link-target"><span>Latviešu</span></a></li><li class="interlanguage-link interwiki-mr mw-list-item"><a href="https://mr.wikipedia.org/wiki/%E0%A4%B8%E0%A4%BF%E0%A4%9F%E0%A5%80,_%E0%A4%AF%E0%A5%81%E0%A4%A8%E0%A4%BF%E0%A4%B5%E0%A5%8D%E0%A4%B9%E0%A4%B0%E0%A5%8D%E0%A4%B8%E0%A4%BF%E0%A4%9F%E0%A5%80_%E0%A4%91%E0%A4%AB_%E0%A4%B2%E0%A4%82%E0%A4%A1%E0%A4%A8" title="सिटी, युनिव्हर्सिटी ऑफ लंडन – Marathi" lang="mr" hreflang="mr" data-title="सिटी, युनिव्हर्सिटी ऑफ लंडन" data-language-autonym="मराठी" data-language-local-name="Marathi" class="interlanguage-link-target"><span>मराठी</span></a></li><li class="interlanguage-link interwiki-arz mw-list-item"><a href="https://arz.wikipedia.org/wiki/%D8%AC%D8%A7%D9%85%D8%B9%D8%A9_%D8%B3%D9%8A%D8%AA%D9%89_%D9%84%D9%86%D8%AF%D9%86" title="جامعة سيتى لندن – Egyptian Arabic" lang="arz" hreflang="arz" data-title="جامعة سيتى لندن" data-language-autonym="مصرى" data-language-local-name="Egyptian Arabic" class="interlanguage-link-target"><span>مصرى</span></a></li><li class="interlanguage-link interwiki-nl mw-list-item"><a href="https://nl.wikipedia.org/wiki/City,_University_of_London" title="City, University of London – Dutch" lang="nl" hreflang="nl" data-title="City, University of London" data-language-autonym="Nederlands" data-language-local-name="Dutch" class="interlanguage-link-target"><span>Nederlands</span></a></li><li class="interlanguage-link interwiki-ja mw-list-item"><a href="https://ja.wikipedia.org/wiki/%E3%83%AD%E3%83%B3%E3%83%89%E3%83%B3%E5%A4%A7%E5%AD%A6%E3%82%B7%E3%83%86%E3%82%A3" title="ロンドン大学シティ – Japanese" lang="ja" hreflang="ja" data-title="ロンドン大学シティ" data-language-autonym="日本語" data-language-local-name="Japanese" class="interlanguage-link-target"><span>日本語</span></a></li><li class="interlanguage-link interwiki-no mw-list-item"><a href="https://no.wikipedia.org/wiki/City_University_London" title="City University London – Norwegian Bokmål" lang="nb" hreflang="nb" data-title="City University London" data-language-autonym="Norsk bokmål" data-language-local-name="Norwegian Bokmål" class="interlanguage-link-target"><span>Norsk bokmål</span></a></li><li class="interlanguage-link interwiki-pnb mw-list-item"><a href="https://pnb.wikipedia.org/wiki/%D8%B3%D9%B9%DB%8C_%DB%8C%D9%88%D9%86%DB%8C%D9%88%D8%B1%D8%B3%D9%B9%DB%8C" title="سٹی یونیورسٹی – Western Punjabi" lang="pnb" hreflang="pnb" data-title="سٹی یونیورسٹی" data-language-autonym="پنجابی" data-language-local-name="Western Punjabi" class="interlanguage-link-target"><span>پنجابی</span></a></li><li class="interlanguage-link interwiki-pl mw-list-item"><a href="https://pl.wikipedia.org/wiki/City_University" title="City University – Polish" lang="pl" hreflang="pl" data-title="City University" data-language-autonym="Polski" data-language-local-name="Polish" class="interlanguage-link-target"><span>Polski</span></a></li><li class="interlanguage-link interwiki-pt mw-list-item"><a href="https://pt.wikipedia.org/wiki/City,_University_of_London" title="City, University of London – Portuguese" lang="pt" hreflang="pt" data-title="City, University of London" data-language-autonym="Português" data-language-local-name="Portuguese" class="interlanguage-link-target"><span>Português</span></a></li><li class="interlanguage-link interwiki-ru mw-list-item"><a href="https://ru.wikipedia.org/wiki/%D0%9B%D0%BE%D0%BD%D0%B4%D0%BE%D0%BD%D1%81%D0%BA%D0%B8%D0%B9_%D0%B3%D0%BE%D1%80%D0%BE%D0%B4%D1%81%D0%BA%D0%BE%D0%B9_%D1%83%D0%BD%D0%B8%D0%B2%D0%B5%D1%80%D1%81%D0%B8%D1%82%D0%B5%D1%82" title="Лондонский городской университет – Russian" lang="ru" hreflang="ru" data-title="Лондонский городской университет" data-language-autonym="Русский" data-language-local-name="Russian" class="interlanguage-link-target"><span>Русский</span></a></li><li class="interlanguage-link interwiki-simple mw-list-item"><a href="https://simple.wikipedia.org/wiki/City,_University_of_London" title="City, University of London – Simple English" lang="en-simple" hreflang="en-simple" data-title="City, University of London" data-language-autonym="Simple English" data-language-local-name="Simple English" class="interlanguage-link-target"><span>Simple English</span></a></li><li class="interlanguage-link interwiki-sv mw-list-item"><a href="https://sv.wikipedia.org/wiki/City,_University_of_London" title="City, University of London – Swedish" lang="sv" hreflang="sv" data-title="City, University of London" data-language-autonym="Svenska" data-language-local-name="Swedish" class="interlanguage-link-target"><span>Svenska</span></a></li><li class="interlanguage-link interwiki-tl mw-list-item"><a href="https://tl.wikipedia.org/wiki/City,_Unibersidad_ng_Londres" title="City, Unibersidad ng Londres – Tagalog" lang="tl" hreflang="tl" data-title="City, Unibersidad ng Londres" data-language-autonym="Tagalog" data-language-local-name="Tagalog" class="interlanguage-link-target"><span>Tagalog</span></a></li><li class="interlanguage-link interwiki-tr mw-list-item"><a href="https://tr.wikipedia.org/wiki/Londra_%C5%9Eehir_%C3%9Cniversitesi" title="Londra Şehir Üniversitesi – Turkish" lang="tr" hreflang="tr" data-title="Londra Şehir Üniversitesi" data-language-autonym="Türkçe" data-language-local-name="Turkish" class="interlanguage-link-target"><span>Türkçe</span></a></li><li class="interlanguage-link interwiki-ur mw-list-item"><a href="https://ur.wikipedia.org/wiki/%D8%B3%D9%B9%DB%8C%D8%8C_%DB%8C%D9%88%D9%86%DB%8C%D9%88%D8%B1%D8%B3%D9%B9%DB%8C_%D8%A2%D9%81_%D9%84%D9%86%D8%AF%D9%86" title="سٹی، یونیورسٹی آف لندن – Urdu" lang="ur" hreflang="ur" data-title="سٹی، یونیورسٹی آف لندن" data-language-autonym="اردو" data-language-local-name="Urdu" class="interlanguage-link-target"><span>اردو</span></a></li><li class="interlanguage-link interwiki-zh-yue mw-list-item"><a href="https://zh-yue.wikipedia.org/wiki/%E5%80%AB%E6%95%A6%E5%A4%A7%E5%AD%B8%E5%9F%8E%E5%B8%82%E5%AD%B8%E9%99%A2" title="倫敦大學城市學院 – Cantonese" lang="yue" hreflang="yue" data-title="倫敦大學城市學院" data-language-autonym="粵語" data-language-local-name="Cantonese" class="interlanguage-link-target"><span>粵語</span></a></li><li class="interlanguage-link interwiki-zh mw-list-item"><a href="https://zh.wikipedia.org/wiki/%E5%80%AB%E6%95%A6%E5%A4%A7%E5%AD%B8%E5%9F%8E%E5%B8%82%E5%AD%B8%E9%99%A2" title="倫敦大學城市學院 – Chinese" lang="zh" hreflang="zh" data-title="倫敦大學城市學院" data-language-autonym="中文" data-language-local-name="Chinese" class="interlanguage-link-target"><span>中文</span></a></li>
			</ul>
			<div class="after-portlet after-portlet-lang"><span class="wb-langlinks-edit wb-langlinks-link"><a href="https://www.wikidata.org/wiki/Special:EntityPage/Q1094046#sitelinks-wikipedia" title="Edit interlanguage links" class="wbc-editpage">Edit links</a></span></div>
		</div>

	</div>
</div>
</header>
				<div class="vector-page-toolbar vector-feature-custom-font-size-clientpref--excluded">
					<div class="vector-page-toolbar-container">
						<div id="left-navigation">
							<nav aria-label="Namespaces">
								
<div id="p-associated-pages" class="vector-menu vector-menu-tabs mw-portlet mw-portlet-associated-pages"  >
	<div class="vector-menu-content">
		
		<ul class="vector-menu-content-list">
			
			<li id="ca-nstab-main" class="selected vector-tab-noicon mw-list-item"><a href="/wiki/City,_University_of_London" title="View the content page [c]" accesskey="c"><span>Article</span></a></li><li id="ca-talk" class="vector-tab-noicon mw-list-item"><a href="/wiki/Talk:City,_University_of_London" rel="discussion" title="Discuss improvements to the content page [t]" accesskey="t"><span>Talk</span></a></li>
		</ul>
		
	</div>
</div>

								
<div id="vector-variants-dropdown" class="vector-dropdown emptyPortlet"  >
	<input type="checkbox" id="vector-variants-dropdown-checkbox" role="button" aria-haspopup="true" data-event-name="ui.dropdown-vector-variants-dropdown" class="vector-dropdown-checkbox " aria-label="Change language variant"   >
	<label id="vector-variants-dropdown-label" for="vector-variants-dropdown-checkbox" class="vector-dropdown-label cdx-button cdx-button--fake-button cdx-button--fake-button--enabled cdx-button--weight-quiet" aria-hidden="true"  ><span class="vector-dropdown-label-text">English</span>
	</label>
	<div class="vector-dropdown-content">


					
<div id="p-variants" class="vector-menu mw-portlet mw-portlet-variants emptyPortlet"  >
	<div class="vector-menu-content">
		
		<ul class="vector-menu-content-list">
			
			
		</ul>
		
	</div>
</div>

				
	</div>
</div>

							</nav>
						</div>
						<div id="right-navigation" class="vector-collapsible">
							<nav aria-label="Views">
								
<div id="p-views" class="vector-menu vector-menu-tabs mw-portlet mw-portlet-views"  >
	<div class="vector-menu-content">
		
		<ul class="vector-menu-content-list">
			
			<li id="ca-view" class="selected vector-tab-noicon mw-list-item"><a href="/wiki/City,_University_of_London"><span>Read</span></a></li><li id="ca-edit" class="vector-tab-noicon mw-list-item"><a href="/w/index.php?title=City,_University_of_London&amp;action=edit" title="Edit this page [e]" accesskey="e"><span>Edit</span></a></li><li id="ca-history" class="vector-tab-noicon mw-list-item"><a href="/w/index.php?title=City,_University_of_London&amp;action=history" title="Past revisions of this page [h]" accesskey="h"><span>View history</span></a></li>
		</ul>
		
	</div>
</div>

							</nav>
				
							<nav class="vector-page-tools-landmark" aria-label="Page tools">
								
<div id="vector-page-tools-dropdown" class="vector-dropdown vector-page-tools-dropdown"  >
	<input type="checkbox" id="vector-page-tools-dropdown-checkbox" role="button" aria-haspopup="true" data-event-name="ui.dropdown-vector-page-tools-dropdown" class="vector-dropdown-checkbox "  aria-label="Tools"  >
	<label id="vector-page-tools-dropdown-label" for="vector-page-tools-dropdown-checkbox" class="vector-dropdown-label cdx-button cdx-button--fake-button cdx-button--fake-button--enabled cdx-button--weight-quiet" aria-hidden="true"  ><span class="vector-dropdown-label-text">Tools</span>
	</label>
	<div class="vector-dropdown-content">


									<div id="vector-page-tools-unpinned-container" class="vector-unpinned-container">
						
<div id="vector-page-tools" class="vector-page-tools vector-pinnable-element">
	<div
	class="vector-pinnable-header vector-page-tools-pinnable-header vector-pinnable-header-unpinned"
	data-feature-name="page-tools-pinned"
	data-pinnable-element-id="vector-page-tools"
	data-pinned-container-id="vector-page-tools-pinned-container"
	data-unpinned-container-id="vector-page-tools-unpinned-container"
>
	<div class="vector-pinnable-header-label">Tools</div>
	<button class="vector-pinnable-header-toggle-button vector-pinnable-header-pin-button" data-event-name="pinnable-header.vector-page-tools.pin">move to sidebar</button>
	<button class="vector-pinnable-header-toggle-button vector-pinnable-header-unpin-button" data-event-name="pinnable-header.vector-page-tools.unpin">hide</button>
</div>

	
<div id="p-cactions" class="vector-menu mw-portlet mw-portlet-cactions emptyPortlet vector-has-collapsible-items"  title="More options" >
	<div class="vector-menu-heading">
		Actions
	</div>
	<div class="vector-menu-content">
		
		<ul class="vector-menu-content-list">
			
			<li id="ca-more-view" class="selected vector-more-collapsible-item mw-list-item"><a href="/wiki/City,_University_of_London"><span>Read</span></a></li><li id="ca-more-edit" class="vector-more-collapsible-item mw-list-item"><a href="/w/index.php?title=City,_University_of_London&amp;action=edit" title="Edit this page [e]" accesskey="e"><span>Edit</span></a></li><li id="ca-more-history" class="vector-more-collapsible-item mw-list-item"><a href="/w/index.php?title=City,_University_of_London&amp;action=history"><span>View history</span></a></li>
		</ul>
		
	</div>
</div>

<div id="p-tb" class="vector-menu mw-portlet mw-portlet-tb"  >
	<div class="vector-menu-heading">
		General
	</div>
	<div class="vector-menu-content">
		
		<ul class="vector-menu-content-list">
			
			<li id="t-whatlinkshere" class="mw-list-item"><a href="/wiki/Special:WhatLinksHere/City,_University_of_London" title="List of all English Wikipedia pages containing links to this page [j]" accesskey="j"><span>What links here</span></a></li><li id="t-recentchangeslinked" class="mw-list-item"><a href="/wiki/Special:RecentChangesLinked/City,_University_of_London" rel="nofollow" title="Recent changes in pages linked from this page [k]" accesskey="k"><span>Related changes</span></a></li><li id="t-upload" class="mw-list-item"><a href="//en.wikipedia.org/wiki/Wikipedia:File_Upload_Wizard" title="Upload files [u]" accesskey="u"><span>Upload file</span></a></li><li id="t-permalink" class="mw-list-item"><a href="/w/index.php?title=City,_University_of_London&amp;oldid=1283711199" title="Permanent link to this revision of this page"><span>Permanent link</span></a></li><li id="t-info" class="mw-list-item"><a href="/w/index.php?title=City,_University_of_London&amp;action=info" title="More information about this page"><span>Page information</span></a></li><li id="t-cite" class="mw-list-item"><a href="/w/index.php?title=Special:CiteThisPage&amp;page=City%2C_University_of_London&amp;id=1283711199&amp;wpFormIdentifier=titleform" title="Information on how to cite this page"><span>Cite this page</span></a></li><li id="t-urlshortener" class="mw-list-item"><a href="/w/index.php?title=Special:UrlShortener&amp;url=https%3A%2F%2Fen.wikipedia.org%2Fwiki%2FCity%2C_University_of_London"><span>Get shortened URL</span></a></li><li id="t-urlshortener-qrcode" class="mw-list-item"><a href="/w/index.php?title=Special:QrCode&amp;url=https%3A%2F%2Fen.wikipedia.org%2Fwiki%2FCity%2C_University_of_London"><span>Download QR code</span></a></li>
		</ul>
		
	</div>
</div>

<div id="p-coll-print_export" class="vector-menu mw-portlet mw-portlet-coll-print_export"  >
	<div class="vector-menu-heading">
		Print/export
	</div>
	<div class="vector-menu-content">
		
		<ul class="vector-menu-content-list">
			
			<li id="coll-download-as-rl" class="mw-list-item"><a href="/w/index.php?title=Special:DownloadAsPdf&amp;page=City%2C_University_of_London&amp;action=show-download-screen" title="Download this page as a PDF file"><span>Download as PDF</span></a></li><li id="t-print" class="mw-list-item"><a href="/w/index.php?title=City,_University_of_London&amp;printable=yes" title="Printable version of this page [p]" accesskey="p"><span>Printable version</span></a></li>
		</ul>
		
	</div>
</div>

<div id="p-wikibase-otherprojects" class="vector-menu mw-portlet mw-portlet-wikibase-otherprojects"  >
	<div class="vector-menu-heading">
		In other projects
	</div>
	<div class="vector-menu-content">
		
		<ul class="vector-menu-content-list">
			
			<li class="wb-otherproject-link wb-otherproject-commons mw-list-item"><a href="https://commons.wikimedia.org/wiki/Category:City,_University_of_London" hreflang="en"><span>Wikimedia Commons</span></a></li><li id="t-wikibase" class="wb-otherproject-link wb-otherproject-wikibase-dataitem mw-list-item"><a href="https://www.wikidata.org/wiki/Special:EntityPage/Q1094046" title="Structured data on this page hosted by Wikidata [g]" accesskey="g"><span>Wikidata item</span></a></li>
		</ul>
		
	</div>
</div>

</div>

									</div>
				
	</div>
</div>

							</nav>
						</div>
					</div>
				</div>
				<div class="vector-column-end no-font-mode-scale">
					<div class="vector-sticky-pinned-container">
						<nav class="vector-page-tools-landmark" aria-label="Page tools">
							<div id="vector-page-tools-pinned-container" class="vector-pinned-container">
				
							</div>
		</nav>
						<nav class="vector-appearance-landmark" aria-label="Appearance">
							<div id="vector-appearance-pinned-container" class="vector-pinned-container">
				<div id="vector-appearance" class="vector-appearance vector-pinnable-element">
	<div
	class="vector-pinnable-header vector-appearance-pinnable-header vector-pinnable-header-pinned"
	data-feature-name="appearance-pinned"
	data-pinnable-element-id="vector-appearance"
	data-pinned-container-id="vector-appearance-pinned-container"
	data-unpinned-container-id="vector-appearance-unpinned-container"
>
	<div class="vector-pinnable-header-label">Appearance</div>
	<button class="vector-pinnable-header-toggle-button vector-pinnable-header-pin-button" data-event-name="pinnable-header.vector-appearance.pin">move to sidebar</button>
	<button class="vector-pinnable-header-toggle-button vector-pinnable-header-unpin-button" data-event-name="pinnable-header.vector-appearance.unpin">hide</button>
</div>


</div>

							</div>
		</nav>
					</div>
				</div>
				<div id="bodyContent" class="vector-body" aria-labelledby="firstHeading" data-mw-ve-target-container>
					<div class="vector-body-before-content">
							<div class="mw-indicators">
		<div id="mw-indicator-coordinates" class="mw-indicator"><div class="mw-parser-output"><span id="coordinates"><a href="/wiki/Geographic_coordinate_system" title="Geographic coordinate system">Coordinates</a>: <style data-mw-deduplicate="TemplateStyles:r1156832818">.mw-parser-output .geo-default,.mw-parser-output .geo-dms,.mw-parser-output .geo-dec{display:inline}.mw-parser-output .geo-nondefault,.mw-parser-output .geo-multi-punct,.mw-parser-output .geo-inline-hidden{display:none}.mw-parser-output .longitude,.mw-parser-output .latitude{white-space:nowrap}</style><span class="plainlinks nourlexpansion"><a class="external text" href="https://geohack.toolforge.org/geohack.php?pagename=City,_University_of_London&amp;params=51.5278_N_0.1023_W_type:edu"><span class="geo-nondefault"><span class="geo-dms" title="Maps, aerial photos, and other data for this location"><span class="latitude">51°31′40″N</span> <span class="longitude">0°06′08″W</span></span></span><span class="geo-multi-punct">&#xfeff; / &#xfeff;</span><span class="geo-default"><span class="geo-dec" title="Maps, aerial photos, and other data for this location">51.5278°N 0.1023°W</span><span style="display:none">&#xfeff; / <span class="geo">51.5278; -0.1023</span></span></span></a></span></span></div></div>
		</div>

						<div id="siteSub" class="noprint">From Wikipedia, the free encyclopedia</div>
					</div>
					<div id="contentSub"><div id="mw-content-subtitle"></div></div>
					
					
					<div id="mw-content-text" class="mw-body-content"><div class="mw-content-ltr mw-parser-output" lang="en" dir="ltr"><div class="shortdescription nomobile noexcerpt noprint searchaux" style="display:none">Former public university in London, United Kingdom</div>
<p class="mw-empty-elt">
</p>
<style data-mw-deduplicate="TemplateStyles:r1251242444">.mw-parser-output .ambox{border:1px solid #a2a9b1;border-left:10px solid #36c;background-color:#fbfbfb;box-sizing:border-box}.mw-parser-output .ambox+link+.ambox,.mw-parser-output .ambox+link+style+.ambox,.mw-parser-output .ambox+link+link+.ambox,.mw-parser-output .ambox+.mw-empty-elt+link+.ambox,.mw-parser-output .ambox+.mw-empty-elt+link+style+.ambox,.mw-parser-output .ambox+.mw-empty-elt+link+link+.ambox{margin-top:-1px}html body.mediawiki .mw-parser-output .ambox.mbox-small-left{margin:4px 1em 4px 0;overflow:hidden;width:238px;border-collapse:collapse;font-size:88%;line-height:1.25em}.mw-parser-output .ambox-speedy{border-left:10px solid #b32424;background-color:#fee7e6}.mw-parser-output .ambox-delete{border-left:10px solid #b32424}.mw-parser-output .ambox-content{border-left:10px solid #f28500}.mw-parser-output .ambox-style{border-left:10px solid #fc3}.mw-parser-output .ambox-move{border-left:10px solid #9932cc}.mw-parser-output .ambox-protection{border-left:10px solid #a2a9b1}.mw-parser-output .ambox .mbox-text{border:none;padding:0.25em 0.5em;width:100%}.mw-parser-output .ambox .mbox-image{border:none;padding:2px 0 2px 0.5em;text-align:center}.mw-parser-output .ambox .mbox-imageright{border:none;padding:2px 0.5em 2px 0;text-align:center}.mw-parser-output .ambox .mbox-empty-cell{border:none;padding:0;width:1px}.mw-parser-output .ambox .mbox-image-div{width:52px}@media(min-width:720px){.mw-parser-output .ambox{margin:0 10%}}@media print{body.ns-0 .mw-parser-output .ambox{display:none!important}}</style><table class="box-Merge_to plainlinks metadata ambox ambox-move" role="presentation"><tbody><tr><td class="mbox-image"><div class="mbox-image-div"><span class="mw-default-size" typeof="mw:File"><span><img alt="" src="//upload.wikimedia.org/wikipedia/commons/thumb/a/aa/Merge-arrow.svg/50px-Merge-arrow.svg.png" decoding="async" width="50" height="20" class="mw-file-element" srcset="//upload.wikimedia.org/wikipedia/commons/thumb/a/aa/Merge-arrow.svg/75px-Merge-arrow.svg.png 1.5x, //upload.wikimedia.org/wikipedia/commons/thumb/a/aa/Merge-arrow.svg/100px-Merge-arrow.svg.png 2x" data-file-width="50" data-file-height="20" /></span></span></div></td><td class="mbox-text"><div class="mbox-text-span">It has been suggested that this article be <a href="/wiki/Wikipedia:Merging" title="Wikipedia:Merging">merged</a> into <i><a href="/wiki/City_St_George%27s,_University_of_London" title="City St George&#39;s, University of London">City St George's, University of London</a></i>. (<a href="/wiki/Talk:City_St_George%27s,_University_of_London#Discussion:_Should_this_article_be_merged?" title="Talk:City St George&#39;s, University of London">Discuss</a>)<small><i> Proposed since March 2025.</i></small></div></td></tr></tbody></table>
<style data-mw-deduplicate="TemplateStyles:r1236090951">.mw-parser-output .hatnote{font-style:italic}.mw-parser-output div.hatnote{padding-left:1.6em;margin-bottom:0.5em}.mw-parser-output .hatnote i{font-style:normal}.mw-parser-output .hatnote+link+.hatnote{margin-top:-0.5em}@media print{body.ns-0 .mw-parser-output .hatnote{display:none!important}}</style><div role="note" class="hatnote navigation-not-searchable">This article is about the former university. For the university after the merge, see <a href="/wiki/City_St_George%27s,_University_of_London" title="City St George&#39;s, University of London">City St George's, University of London</a>.</div><style data-mw-deduplicate="TemplateStyles:r1289430074">.mw-parser-output .infobox-subbox{padding:0;border:none;margin:-3px;width:auto;min-width:100%;font-size:100%;clear:none;float:none;background-color:transparent}.mw-parser-output .infobox-3cols-child{margin:auto}.mw-parser-output .infobox .navbar{font-size:100%}@media screen{html.skin-theme-clientpref-night .mw-parser-output .infobox-full-data:not(.notheme)>div:not(.notheme)[style]{background:#1f1f23!important;color:#f8f9fa}}@media screen and (prefers-color-scheme:dark){html.skin-theme-clientpref-os .mw-parser-output .infobox-full-data:not(.notheme) div:not(.notheme){background:#1f1f23!important;color:#f8f9fa}}@media(min-width:640px){body.skin--responsive .mw-parser-output .infobox-table{display:table!important}body.skin--responsive .mw-parser-output .infobox-table>caption{display:table-caption!important}body.skin--responsive .mw-parser-output .infobox-table>tbody{display:table-row-group}body.skin--responsive .mw-parser-output .infobox-table th,body.skin--responsive .mw-parser-output .infobox-table td{padding-left:inherit;padding-right:inherit}}</style><table class="infobox vcard"><caption class="infobox-title fn org">City, University of London</caption><tbody><tr><td colspan="2" class="infobox-image"><span class="mw-default-size" typeof="mw:File/Frameless"><a href="/wiki/File:Arms_of_City,_University_of_London.svg" class="mw-file-description"><img src="//upload.wikimedia.org/wikipedia/commons/thumb/0/01/Arms_of_City%2C_University_of_London.svg/250px-Arms_of_City%2C_University_of_London.svg.png" decoding="async" width="180" height="241" class="mw-file-element" srcset="//upload.wikimedia.org/wikipedia/commons/thumb/0/01/Arms_of_City%2C_University_of_London.svg/330px-Arms_of_City%2C_University_of_London.svg.png 1.5x, //upload.wikimedia.org/wikipedia/commons/thumb/0/01/Arms_of_City%2C_University_of_London.svg/500px-Arms_of_City%2C_University_of_London.svg.png 2x" data-file-width="750" data-file-height="1003" /></a></span><div class="infobox-caption"><a href="/wiki/Coat_of_arms" title="Coat of arms">Coat of arms</a> of the university</div></td></tr><tr><th scope="row" class="infobox-label" style="padding-right:0.65em;">Motto</th><td class="infobox-data">To Serve Mankind</td></tr><tr><th scope="row" class="infobox-label" style="padding-right:0.65em;">Type</th><td class="infobox-data"><a href="/wiki/Public_university" title="Public university">Public</a> <a href="/wiki/Research_university" title="Research university">research university</a></td></tr><tr><th scope="row" class="infobox-label" style="padding-right:0.65em;">Established</th><td class="infobox-data">1852 – <a href="/wiki/Inns_of_Court_School_of_Law" class="mw-redirect" title="Inns of Court School of Law">Inns of Court School of Law</a><br />1894 – Northampton Institute<br />1966 – gained <a href="/wiki/Universities_in_the_United_Kingdom" title="Universities in the United Kingdom">university status</a> by <a href="/wiki/Royal_charter" title="Royal charter">royal charter</a><br />2016 – constituent college of <a href="/wiki/University_of_London" title="University of London">University of London</a><br />2024 – merged with <a href="/wiki/St_George%27s,_University_of_London" title="St George&#39;s, University of London">St George's, University of London</a></td></tr><tr><th scope="row" class="infobox-label" style="padding-right:0.65em;"><a href="/wiki/Financial_endowment" title="Financial endowment">Endowment</a></th><td class="infobox-data"><a href="/wiki/Pound_sterling" title="Pound sterling">£</a>22.0&#160;million (2023–24)<sup id="cite_ref-City_Financial_Statement_23/24_1-0" class="reference"><a href="#cite_note-City_Financial_Statement_23/24-1"><span class="cite-bracket">&#91;</span>1<span class="cite-bracket">&#93;</span></a></sup></td></tr><tr><th scope="row" class="infobox-label" style="padding-right:0.65em;">Budget</th><td class="infobox-data"><a href="/wiki/Pounds_sterling" class="mw-redirect" title="Pounds sterling">£</a>301.7 million (2023–24)<sup id="cite_ref-City_Financial_Statement_23/24_1-1" class="reference"><a href="#cite_note-City_Financial_Statement_23/24-1"><span class="cite-bracket">&#91;</span>1<span class="cite-bracket">&#93;</span></a></sup></td></tr><tr><th scope="row" class="infobox-label" style="padding-right:0.65em;"><a href="/wiki/Chancellor_(education)" title="Chancellor (education)">Chancellor</a></th><td class="infobox-data"><a href="/wiki/Anne,_Princess_Royal" title="Anne, Princess Royal">The Princess Royal</a><br />(as Chancellor of the <a href="/wiki/University_of_London" title="University of London">University of London</a>)</td></tr><tr><th scope="row" class="infobox-label" style="padding-right:0.65em;"><a href="/wiki/University_president" class="mw-redirect" title="University president">President</a></th><td class="infobox-data">Sir <a href="/wiki/Anthony_Finkelstein" title="Anthony Finkelstein">Anthony Finkelstein</a></td></tr><tr><th scope="row" class="infobox-label" style="padding-right:0.65em;"><a href="/wiki/Rector_(academia)" title="Rector (academia)">Rector</a></th><td class="infobox-data"><a href="/wiki/Lord_Mayor_of_the_City_of_London" class="mw-redirect" title="Lord Mayor of the City of London">Lord Mayor of the City of London</a> (<i><a href="/wiki/Ex_officio" class="mw-redirect" title="Ex officio">ex officio</a></i>)</td></tr><tr><th scope="row" class="infobox-label" style="padding-right:0.65em;">Students</th><td class="infobox-data">21,735 (2022/23)<sup id="cite_ref-HESAWhere2024_2-0" class="reference"><a href="#cite_note-HESAWhere2024-2"><span class="cite-bracket">&#91;</span>2<span class="cite-bracket">&#93;</span></a></sup></td></tr><tr><th scope="row" class="infobox-label" style="padding-right:0.65em;"><a href="/wiki/Undergraduate_education" title="Undergraduate education">Undergraduates</a></th><td class="infobox-data">13,590 (2022/23)<sup id="cite_ref-HESAWhere2024_2-1" class="reference"><a href="#cite_note-HESAWhere2024-2"><span class="cite-bracket">&#91;</span>2<span class="cite-bracket">&#93;</span></a></sup></td></tr><tr><th scope="row" class="infobox-label" style="padding-right:0.65em;"><a href="/wiki/Postgraduate_education" title="Postgraduate education">Postgraduates</a></th><td class="infobox-data">8,145 (2022/23)<sup id="cite_ref-HESAWhere2024_2-2" class="reference"><a href="#cite_note-HESAWhere2024-2"><span class="cite-bracket">&#91;</span>2<span class="cite-bracket">&#93;</span></a></sup></td></tr><tr><th scope="row" class="infobox-label" style="padding-right:0.65em;">Location</th><td class="infobox-data adr"><div style="display:inline" class="locality"><a href="/wiki/London" title="London">London</a></div>, <div style="display:inline" class="country-name">United Kingdom</div><br /><span class="geo-inline"><style data-mw-deduplicate="TemplateStyles:r1156832818">.mw-parser-output .geo-default,.mw-parser-output .geo-dms,.mw-parser-output .geo-dec{display:inline}.mw-parser-output .geo-nondefault,.mw-parser-output .geo-multi-punct,.mw-parser-output .geo-inline-hidden{display:none}.mw-parser-output .longitude,.mw-parser-output .latitude{white-space:nowrap}</style><span class="plainlinks nourlexpansion"><a class="external text" href="https://geohack.toolforge.org/geohack.php?pagename=City,_University_of_London&amp;params=51.5278_N_0.1023_W_type:edu"><span class="geo-nondefault"><span class="geo-dms" title="Maps, aerial photos, and other data for this location"><span class="latitude">51°31′40″N</span> <span class="longitude">0°06′08″W</span></span></span><span class="geo-multi-punct">&#xfeff; / &#xfeff;</span><span class="geo-default"><span class="geo-dec" title="Maps, aerial photos, and other data for this location">51.5278°N 0.1023°W</span><span style="display:none">&#xfeff; / <span class="geo">51.5278; -0.1023</span></span></span></a></span></span></td></tr><tr><th scope="row" class="infobox-label" style="padding-right:0.65em;">Campus</th><td class="infobox-data">Urban</td></tr><tr><th scope="row" class="infobox-label" style="padding-right:0.65em;"><a href="/wiki/School_colors" title="School colors">Colours</a></th><td class="infobox-data">Red and white <table role="presentation" style="width:100%; border-collapse:collapse;">
<tbody><tr style="border: 1px solid #a2a9b1; background-color: #f8f9fa; color:inherit; padding: 5px; font-size: 95%;height:1.75em"><td style="background: #e41f13; color:inherit;">&#8201;</td>
<td style="background: #e41f13; color:inherit;">&#8201;</td>
<td style="background: #e41f13; color:inherit;">&#8201;</td>
<td style="background: #e41f13; color:inherit;">&#8201;</td>
<td style="background: #e41f13; color:inherit;">&#8201;</td><td style="background: #FFF; color:inherit;">&#8201;</td>
<td style="background: #FFF; color:inherit;">&#8201;</td>
<td style="background: #FFF; color:inherit;">&#8201;</td>
<td style="background: #FFF; color:inherit;">&#8201;</td>
<td style="background: #FFF; color:inherit;">&#8201;</td><td style="background: #e41f13; color:inherit;">&#8201;</td>
<td style="background: #e41f13; color:inherit;">&#8201;</td>
<td style="background: #e41f13; color:inherit;">&#8201;</td>
<td style="background: #e41f13; color:inherit;">&#8201;</td>
<td style="background: #e41f13; color:inherit;">&#8201;</td></tr>
      </tbody></table></td></tr><tr><th scope="row" class="infobox-label" style="padding-right:0.65em;">Affiliations</th><td class="infobox-data"><a href="/wiki/University_of_London" title="University of London">University of London</a><br /><a href="/wiki/Association_of_MBAs" title="Association of MBAs">Association of MBAs</a><br /><a href="/wiki/European_Quality_Improvement_System" class="mw-redirect" title="European Quality Improvement System">EQUIS</a><br /><a href="/wiki/Universities_UK" title="Universities UK">Universities UK</a></td></tr><tr><th scope="row" class="infobox-label" style="padding-right:0.65em;">Website</th><td class="infobox-data"><span class="url"><a rel="nofollow" class="external text" href="https://www.citystgeorges.ac.uk/">www<wbr />.citystgeorges<wbr />.ac<wbr />.uk</a></span></td></tr><tr><td colspan="2" class="infobox-full-data"><span class="mw-default-size" typeof="mw:File/Frameless"><a href="/wiki/File:City,_University_of_London_Logo.svg" class="mw-file-description"><img src="//upload.wikimedia.org/wikipedia/en/thumb/8/8a/City%2C_University_of_London_Logo.svg/250px-City%2C_University_of_London_Logo.svg.png" decoding="async" width="180" height="106" class="mw-file-element" srcset="//upload.wikimedia.org/wikipedia/en/thumb/8/8a/City%2C_University_of_London_Logo.svg/330px-City%2C_University_of_London_Logo.svg.png 1.5x, //upload.wikimedia.org/wikipedia/en/thumb/8/8a/City%2C_University_of_London_Logo.svg/500px-City%2C_University_of_London_Logo.svg.png 2x" data-file-width="512" data-file-height="302" /></a></span></td></tr></tbody></table>
<p><b>City, University of London</b> was a <a href="/wiki/Public_university" title="Public university">public university</a> from 1966 to 2024 in <a href="/wiki/London" title="London">London</a>, <a href="/wiki/England" title="England">England</a>. It merged with <a href="/wiki/St_George%27s,_University_of_London" title="St George&#39;s, University of London">St George's, University of London</a> to form <a href="/wiki/City_St_George%27s,_University_of_London" title="City St George&#39;s, University of London">City St George's, University of London</a> in August 2024.<sup id="cite_ref-3" class="reference"><a href="#cite_note-3"><span class="cite-bracket">&#91;</span>3<span class="cite-bracket">&#93;</span></a></sup> The names "City, University of London" and "St George’s, University of London" will provisionally continue as trading names until March 2025.<sup id="cite_ref-4" class="reference"><a href="#cite_note-4"><span class="cite-bracket">&#91;</span>4<span class="cite-bracket">&#93;</span></a></sup>
</p><p>Originally founded in 1894 as the <b>Northampton Institute</b>, it officially became a university when <b>The City University</b> was created by <a href="/wiki/Royal_charter" title="Royal charter">royal charter</a> in 1966.<sup id="cite_ref-5" class="reference"><a href="#cite_note-5"><span class="cite-bracket">&#91;</span>5<span class="cite-bracket">&#93;</span></a></sup> The <a href="/wiki/Inns_of_Court_School_of_Law" class="mw-redirect" title="Inns of Court School of Law">Inns of Court School of Law</a>, which merged with City in 2001, was established in 1852, making it the university's oldest constituent part.<sup id="cite_ref-6" class="reference"><a href="#cite_note-6"><span class="cite-bracket">&#91;</span>6<span class="cite-bracket">&#93;</span></a></sup> City joined the federal University of London on 1 September 2016, becoming part of the eighteen colleges and ten research institutes that then made up that university.<sup id="cite_ref-Grove_7-0" class="reference"><a href="#cite_note-Grove-7"><span class="cite-bracket">&#91;</span>7<span class="cite-bracket">&#93;</span></a></sup>
</p><p>City has strong links with the <a href="/wiki/City_of_London" title="City of London">City of London</a>, and the <a href="/wiki/Lord_Mayor_of_London" title="Lord Mayor of London">Lord Mayor of London</a> serves as the university's rector.<sup id="cite_ref-8" class="reference"><a href="#cite_note-8"><span class="cite-bracket">&#91;</span>8<span class="cite-bracket">&#93;</span></a></sup><sup id="cite_ref-9" class="reference"><a href="#cite_note-9"><span class="cite-bracket">&#91;</span>9<span class="cite-bracket">&#93;</span></a></sup> The university has its main campus in <a href="/wiki/Central_London" title="Central London">Central London</a> in the <a href="/wiki/London_Borough_of_Islington" title="London Borough of Islington">London Borough of Islington</a>, with additional campuses in Islington, the <a href="/wiki/City_of_London" title="City of London">City of London</a>, the <a href="/wiki/West_End_of_London" title="West End of London">West End</a> and <a href="/wiki/East_End_of_London" title="East End of London">East End</a>. It is organised into six schools, within which there are around forty academic departments and centres,<sup id="cite_ref-10" class="reference"><a href="#cite_note-10"><span class="cite-bracket">&#91;</span>10<span class="cite-bracket">&#93;</span></a></sup> including the <a href="/wiki/Department_of_Journalism,_City_University" title="Department of Journalism, City University">Department of Journalism</a>, <a href="/wiki/Bayes_Business_School" title="Bayes Business School">Bayes Business School</a> (formerly Cass Business School), and <a href="/wiki/City_Law_School" title="City Law School">City Law School</a> which incorporates the Inns of Court School of Law.<sup id="cite_ref-11" class="reference"><a href="#cite_note-11"><span class="cite-bracket">&#91;</span>11<span class="cite-bracket">&#93;</span></a></sup> The annual income of the institution for 2021–22 was £262.1&#160;million, of which £12.9&#160;million was from research grants and contracts, with an expenditure of £328.2&#160;million.<sup id="cite_ref-City_Financial_Statement_21/22_12-0" class="reference"><a href="#cite_note-City_Financial_Statement_21/22-12"><span class="cite-bracket">&#91;</span>12<span class="cite-bracket">&#93;</span></a></sup>
</p><p>City is a founding member of the WC2 University Network which developed for collaboration between leading universities of the heart of major world cities particularly to address cultural, environmental and political issues of common interest to world cities and their universities.<sup id="cite_ref-:1_13-0" class="reference"><a href="#cite_note-:1-13"><span class="cite-bracket">&#91;</span>13<span class="cite-bracket">&#93;</span></a></sup> The university is a member of the <a href="/wiki/Association_of_MBAs" title="Association of MBAs">Association of MBAs</a>, <a href="/wiki/European_Quality_Improvement_System" class="mw-redirect" title="European Quality Improvement System">EQUIS</a> and <a href="/wiki/Universities_UK" title="Universities UK">Universities UK</a>. Alumni of City include <a href="/wiki/Mahatma_Gandhi" title="Mahatma Gandhi">Mahatma Gandhi</a>, <a href="/wiki/Muhammad_Ali_Jinnah" title="Muhammad Ali Jinnah">Muhammad Ali Jinnah</a>,<sup id="cite_ref-14" class="reference"><a href="#cite_note-14"><span class="cite-bracket">&#91;</span>14<span class="cite-bracket">&#93;</span></a></sup> members of <a href="/wiki/Parliament_of_the_United_Kingdom" title="Parliament of the United Kingdom">Parliament of the United Kingdom</a>, <a href="/wiki/Governor" title="Governor">governors</a>, politicians and CEOs.
</p>
<style data-mw-deduplicate="TemplateStyles:r886046785">.mw-parser-output .toclimit-2 .toclevel-1 ul,.mw-parser-output .toclimit-3 .toclevel-2 ul,.mw-parser-output .toclimit-4 .toclevel-3 ul,.mw-parser-output .toclimit-5 .toclevel-4 ul,.mw-parser-output .toclimit-6 .toclevel-5 ul,.mw-parser-output .toclimit-7 .toclevel-6 ul{display:none}</style><div class="toclimit-3"><meta property="mw:PageProp/toc" /></div>
<div class="mw-heading mw-heading2"><h2 id="History">History</h2><span class="mw-editsection"><span class="mw-editsection-bracket">[</span><a href="/w/index.php?title=City,_University_of_London&amp;action=edit&amp;section=1" title="Edit section: History"><span>edit</span></a><span class="mw-editsection-bracket">]</span></span></div>
<div class="mw-heading mw-heading3"><h3 id="Origins">Origins</h3><span class="mw-editsection"><span class="mw-editsection-bracket">[</span><a href="/w/index.php?title=City,_University_of_London&amp;action=edit&amp;section=2" title="Edit section: Origins"><span>edit</span></a><span class="mw-editsection-bracket">]</span></span></div>
<figure class="mw-halign-right" typeof="mw:File/Thumb"><a href="/wiki/File:Northampton_squ.jpg" class="mw-file-description"><img src="//upload.wikimedia.org/wikipedia/commons/thumb/c/c7/Northampton_squ.jpg/250px-Northampton_squ.jpg" decoding="async" width="220" height="165" class="mw-file-element" srcset="//upload.wikimedia.org/wikipedia/commons/thumb/c/c7/Northampton_squ.jpg/330px-Northampton_squ.jpg 1.5x, //upload.wikimedia.org/wikipedia/commons/thumb/c/c7/Northampton_squ.jpg/440px-Northampton_squ.jpg 2x" data-file-width="496" data-file-height="373" /></a><figcaption>Northampton Square in front of the main university building</figcaption></figure>
<p>City traces its origin to the Northampton Institute and <a href="/wiki/The_City_Law_School" class="mw-redirect" title="The City Law School">the City Law School</a> (established in 1852). The first was named after the <a href="/wiki/William_Compton,_4th_Marquess_of_Northampton" title="William Compton, 4th Marquess of Northampton">Marquess of Northampton</a> who donated the land on which the institute was built, between <a href="/wiki/Northampton_Square" title="Northampton Square">Northampton Square</a> and St John Street in <a href="/wiki/Islington" title="Islington">Islington</a>. The institute was established to provide for the education and welfare of the local population. It was constituted under the City of London Parochial Charities Act (1883), with the objective of "the promotion of the industrial skill, general knowledge, health and well-being of young men and women belonging to the poorer classes".<sup id="cite_ref-CityHistory_15-0" class="reference"><a href="#cite_note-CityHistory-15"><span class="cite-bracket">&#91;</span>15<span class="cite-bracket">&#93;</span></a></sup>
</p><p>Northampton Polytechnic Institute was an <a href="/wiki/Institute_of_technology" title="Institute of technology">institute of technology</a> in <a href="/wiki/Clerkenwell" title="Clerkenwell">Clerkenwell</a>, London, founded in 1894. Its first Principal was <a href="/wiki/Robert_Mullineux_Walmsley" title="Robert Mullineux Walmsley">Robert Mullineux Walmsley</a>.<sup id="cite_ref-16" class="reference"><a href="#cite_note-16"><span class="cite-bracket">&#91;</span>16<span class="cite-bracket">&#93;</span></a></sup>
</p><p>Alumni include <a href="/wiki/Colin_Cherry" title="Colin Cherry">Colin Cherry</a>, <a href="/wiki/Stuart_Davies_(engineer)" title="Stuart Davies (engineer)">Stuart Davies</a> and <a href="/wiki/Anthony_Hunt" title="Anthony Hunt">Anthony Hunt</a>.<sup id="cite_ref-17" class="reference"><a href="#cite_note-17"><span class="cite-bracket">&#91;</span>17<span class="cite-bracket">&#93;</span></a></sup> <a href="/wiki/Arthur_George_Cocksedge" class="mw-redirect" title="Arthur George Cocksedge">Arthur George Cocksedge</a>, a British <a href="/wiki/Gymnast" class="mw-redirect" title="Gymnast">gymnast</a> who competed in the <a href="/wiki/1920_Summer_Olympics" title="1920 Summer Olympics">1920 Summer Olympics</a>, was a member of the Northampton Polytechnic Institute's Gymnastics Club and was Champion of the United Kingdom in 1920. In 1937 <a href="/wiki/Maurice_Dennis" title="Maurice Dennis">Maurice Dennis</a> of the (Northampton Polytechnic ABC) was the 1937 <a href="/wiki/ABA_Middleweight_Champion" class="mw-redirect" title="ABA Middleweight Champion">ABA Middleweight Champion</a>. <a href="/wiki/Frederick_Handley_Page" title="Frederick Handley Page">Frederick Handley Page</a> was a lecturer in <a href="/wiki/Aeronautics" title="Aeronautics">aeronautics</a> at the institute. The <a href="/wiki/Handley_Page_Type_A" title="Handley Page Type A">Handley Page Type A</a>, the first powered aircraft designed and built by him, ended up as an instructional airframe at the school. The novelist <a href="/wiki/Eric_Ambler" title="Eric Ambler">Eric Ambler</a> studied engineering at the institute.<sup class="noprint Inline-Template Template-Fact" style="white-space:nowrap;">&#91;<i><a href="/wiki/Wikipedia:Citation_needed" title="Wikipedia:Citation needed"><span title="This claim needs references to reliable sources. (October 2023)">citation needed</span></a></i>&#93;</sup>
</p><p>The six original departments at the institute were Applied Physics and Electrical Engineering; Artistic Crafts; Domestic Economy and Women's Trades; Electro-Chemistry; <a href="/wiki/Horology" class="mw-redirect" title="Horology">Horology</a>; and Mechanical Engineering and Metal Trades.
</p>
<div class="mw-heading mw-heading3"><h3 id="20th_century">20th century</h3><span class="mw-editsection"><span class="mw-editsection-bracket">[</span><a href="/w/index.php?title=City,_University_of_London&amp;action=edit&amp;section=3" title="Edit section: 20th century"><span>edit</span></a><span class="mw-editsection-bracket">]</span></span></div>
<p>A separate technical <a href="/wiki/Optics" title="Optics">optics</a> department was established in 1903–04. In 1909, the first students qualified for <a href="/wiki/University_of_London" title="University of London">University of London</a> BSc degrees in engineering as internal students.<sup id="cite_ref-CityHistory_15-1" class="reference"><a href="#cite_note-CityHistory-15"><span class="cite-bracket">&#91;</span>15<span class="cite-bracket">&#93;</span></a></sup> The Institute had been involved in <a href="/wiki/Aeronautics" title="Aeronautics">aeronautics</a> education since that year, and the School of Engineering and Mathematical Sciences celebrated the centenary of aeronautics at City in 2009.<sup id="cite_ref-18" class="reference"><a href="#cite_note-18"><span class="cite-bracket">&#91;</span>18<span class="cite-bracket">&#93;</span></a></sup> The institute was used for the <a href="/wiki/1908_Olympic_Games" class="mw-redirect" title="1908 Olympic Games">1908 Olympic Games</a>;<sup id="cite_ref-CityHistory_15-2" class="reference"><a href="#cite_note-CityHistory-15"><span class="cite-bracket">&#91;</span>15<span class="cite-bracket">&#93;</span></a></sup> <a href="/wiki/Boxing_at_the_1908_Summer_Olympics" title="Boxing at the 1908 Summer Olympics">boxing</a> took place there.<sup id="cite_ref-19" class="reference"><a href="#cite_note-19"><span class="cite-bracket">&#91;</span>19<span class="cite-bracket">&#93;</span></a></sup>
</p><p>In 1957, the institute was designated a "<a href="/wiki/College_of_Advanced_Technology_(United_Kingdom)" class="mw-redirect" title="College of Advanced Technology (United Kingdom)">College of Advanced Technology</a>".<sup id="cite_ref-CityHistory_15-3" class="reference"><a href="#cite_note-CityHistory-15"><span class="cite-bracket">&#91;</span>15<span class="cite-bracket">&#93;</span></a></sup>
</p><p>The institute's involvement in <a href="/wiki/Information_science" title="Information science">information science</a> began in 1961, with the introduction of a course on "Collecting and Communicating Scientific Knowledge". City received its <a href="/wiki/Royal_charter" title="Royal charter">royal charter</a> in 1966, becoming "The City University" to reflect the institution's close links with the City of London.<sup id="cite_ref-20" class="reference"><a href="#cite_note-20"><span class="cite-bracket">&#91;</span>20<span class="cite-bracket">&#93;</span></a></sup> The <a href="/wiki/Apollo_15" title="Apollo 15">Apollo 15</a> astronauts visited City in 1971, and presented the Vice-Chancellor, Tait, with a piece of <a href="/wiki/Heat_shield" title="Heat shield">heat shield</a> from the Apollo 15 <a href="/wiki/Rocket" title="Rocket">rocket</a>.<sup id="cite_ref-21" class="reference"><a href="#cite_note-21"><span class="cite-bracket">&#91;</span>21<span class="cite-bracket">&#93;</span></a></sup>
</p><p>In October 1995, it was announced that City University would merge with both the <a href="/wiki/St_Bartholomew_School_of_Nursing_%26_Midwifery" class="mw-redirect" title="St Bartholomew School of Nursing &amp; Midwifery">St Bartholomew School of Nursing &amp; Midwifery</a> and the Charterhouse College of Radiography, doubling the number of students in City's Institute of Health Sciences to around 2,500.<sup id="cite_ref-22" class="reference"><a href="#cite_note-22"><span class="cite-bracket">&#91;</span>22<span class="cite-bracket">&#93;</span></a></sup>
</p>
<div class="mw-heading mw-heading3"><h3 id="21st_century">21st century</h3><span class="mw-editsection"><span class="mw-editsection-bracket">[</span><a href="/w/index.php?title=City,_University_of_London&amp;action=edit&amp;section=4" title="Edit section: 21st century"><span>edit</span></a><span class="mw-editsection-bracket">]</span></span></div>
<p>The university formed a strategic alliance with <a href="/wiki/Queen_Mary,_University_of_London" class="mw-redirect" title="Queen Mary, University of London">Queen Mary, University of London</a>, in April 2001.<sup id="cite_ref-23" class="reference"><a href="#cite_note-23"><span class="cite-bracket">&#91;</span>23<span class="cite-bracket">&#93;</span></a></sup> In May 2001, a fire in the college building gutted the fourth-floor offices and roof.<sup id="cite_ref-24" class="reference"><a href="#cite_note-24"><span class="cite-bracket">&#91;</span>24<span class="cite-bracket">&#93;</span></a></sup> In August 2001 City and the <a href="/wiki/Inns_of_Court_School_of_Law" class="mw-redirect" title="Inns of Court School of Law">Inns of Court School of Law</a> agreed to merge.<sup id="cite_ref-25" class="reference"><a href="#cite_note-25"><span class="cite-bracket">&#91;</span>25<span class="cite-bracket">&#93;</span></a></sup> Following a donation from <a href="/wiki/Sir_John_Cass%27s_Foundation" class="mw-redirect" title="Sir John Cass&#39;s Foundation">Sir John Cass's Foundation</a>, a multimillion-pound building was built at 106 Bunhill Row for the Business School.<sup id="cite_ref-26" class="reference"><a href="#cite_note-26"><span class="cite-bracket">&#91;</span>26<span class="cite-bracket">&#93;</span></a></sup>
</p>
<figure class="mw-halign-right" typeof="mw:File/Thumb"><a href="/wiki/File:CollegeBuilding.jpg" class="mw-file-description"><img src="//upload.wikimedia.org/wikipedia/commons/thumb/3/31/CollegeBuilding.jpg/250px-CollegeBuilding.jpg" decoding="async" width="220" height="168" class="mw-file-element" srcset="//upload.wikimedia.org/wikipedia/commons/thumb/3/31/CollegeBuilding.jpg/330px-CollegeBuilding.jpg 1.5x, //upload.wikimedia.org/wikipedia/commons/thumb/3/31/CollegeBuilding.jpg/440px-CollegeBuilding.jpg 2x" data-file-width="450" data-file-height="344" /></a><figcaption>The Grade II listed College Building</figcaption></figure>
<p>A new £23 million building to house the School of Social Sciences and the Department of Language and Communication Science was opened in 2004. The reconstruction and redevelopment of the university's <a href="/wiki/Grade_II_listed" class="mw-redirect" title="Grade II listed">Grade II listed</a> college building (following the fire in 2001) was completed in July 2006.
</p><p>In 2007 the School of Arts received a £10m building refurbishment. A new students' union venue opened in October 2008 called "TEN squared", which provides a hub for students to socialise in during the day and hosts a wide range of evening entertainment including club nights, society events and quiz nights.
</p><p>In January 2010, premises were shared with the <a href="/wiki/University_of_East_Anglia" title="University of East Anglia">University of East Anglia</a> (UEA) London, following City's partnership with <a href="/wiki/INTO_University_Partnerships" title="INTO University Partnerships">INTO University Partnerships</a>. Since then City has resumed its own International Foundation Programme to prepare students for their pre-university year. City was ranked among the top 30 higher education institutions in the UK by the <i>Times Higher Education Table of Tables</i>.<sup id="cite_ref-27" class="reference"><a href="#cite_note-27"><span class="cite-bracket">&#91;</span>27<span class="cite-bracket">&#93;</span></a></sup>
</p><p>In April 2011, it was announced that the current halls of residence and Saddler's Sports Centre will be closed and demolished for rebuilding in June 2011. The new student halls and sports facility, now known as CitySport, opened in 2015.
</p><p>In September 2016 The City University became a member institution of the federal <a href="/wiki/University_of_London" title="University of London">University of London</a><sup id="cite_ref-Grove_7-1" class="reference"><a href="#cite_note-Grove-7"><span class="cite-bracket">&#91;</span>7<span class="cite-bracket">&#93;</span></a></sup> and changed its name to City, University of London.
</p><p>In 2023, a merger was proposed between City and <a href="/wiki/St_George%27s,_University_of_London" title="St George&#39;s, University of London">St George's, University of London</a>.<sup id="cite_ref-28" class="reference"><a href="#cite_note-28"><span class="cite-bracket">&#91;</span>28<span class="cite-bracket">&#93;</span></a></sup> On 1 August 2024, City merged with <a href="/wiki/St_George%27s,_University_of_London" title="St George&#39;s, University of London">St George's, University of London</a> to form <a href="/wiki/City_St_George%27s,_University_of_London" title="City St George&#39;s, University of London">City St George's, University of London</a>.<sup id="cite_ref-29" class="reference"><a href="#cite_note-29"><span class="cite-bracket">&#91;</span>29<span class="cite-bracket">&#93;</span></a></sup>
</p>
<div class="mw-heading mw-heading2"><h2 id="Campus">Campus</h2><span class="mw-editsection"><span class="mw-editsection-bracket">[</span><a href="/w/index.php?title=City,_University_of_London&amp;action=edit&amp;section=5" title="Edit section: Campus"><span>edit</span></a><span class="mw-editsection-bracket">]</span></span></div>
<style data-mw-deduplicate="TemplateStyles:r1291884410">.mw-parser-output .locmap .od{position:absolute}.mw-parser-output .locmap .id{position:absolute;line-height:0}.mw-parser-output .locmap .l0{font-size:0;position:absolute}.mw-parser-output .locmap .pv{line-height:110%;position:absolute;text-align:center}.mw-parser-output .locmap .pl{line-height:110%;position:absolute;top:-0.75em;text-align:right}.mw-parser-output .locmap .pr{line-height:110%;position:absolute;top:-0.75em;text-align:left}.mw-parser-output .locmap .pv>div{display:inline;padding:1px}.mw-parser-output .locmap .pl>div{display:inline;padding:1px;float:right}.mw-parser-output .locmap .pr>div{display:inline;padding:1px;float:left}@media screen{html.skin-theme-clientpref-night .mw-parser-output .od,html.skin-theme-clientpref-night .mw-parser-output .od .pv>div,html.skin-theme-clientpref-night .mw-parser-output .od .pl>div,html.skin-theme-clientpref-night .mw-parser-output .od .pr>div{background:#fff!important;color:#000!important}html.skin-theme-clientpref-night .mw-parser-output .locmap img{filter:grayscale(0.6)}html.skin-theme-clientpref-night .mw-parser-output .infobox-full-data .locmap div{background:transparent!important}}@media screen and (prefers-color-scheme:dark){html.skin-theme-clientpref-os .mw-parser-output .locmap img{filter:grayscale(0.6)}html.skin-theme-clientpref-os .mw-parser-output .od,html.skin-theme-clientpref-os .mw-parser-output .od .pv>div,html.skin-theme-clientpref-os .mw-parser-output .od .pl>div,html.skin-theme-clientpref-os .mw-parser-output .od .pr>div{background:white!important;color:#000!important}html.skin-theme-clientpref-os .mw-parser-output .infobox-full-data .locmap div{background:transparent!important}}</style><div class="locmap noviewer noresize thumb tright"><div class="thumbinner" style="width:242px"><div style="position:relative;width:240px;border:1px solid lightgray"><span class="notpageimage" typeof="mw:File"><a href="/wiki/File:Open_street_map_central_london.svg" class="mw-file-description" title="City, University of London is located in Central London"><img alt="City, University of London is located in Central London" src="//upload.wikimedia.org/wikipedia/commons/thumb/4/4b/Open_street_map_central_london.svg/250px-Open_street_map_central_london.svg.png" decoding="async" width="240" height="152" class="mw-file-element" srcset="//upload.wikimedia.org/wikipedia/commons/thumb/4/4b/Open_street_map_central_london.svg/500px-Open_street_map_central_london.svg.png 1.5x" data-file-width="1193" data-file-height="754" /></a></span><div class="od notheme" style="top:18.417%;left:74.097%;font-size:91%"><div class="id" style="left:-4px;top:-4px"><span class="notpageimage" typeof="mw:File"><span title="City, University of London"><img alt="City, University of London" src="//upload.wikimedia.org/wikipedia/commons/thumb/0/0c/Red_pog.svg/20px-Red_pog.svg.png" decoding="async" width="8" height="8" class="mw-file-element" data-file-width="64" data-file-height="64" /></span></span></div></div></div><div class="thumbcaption"><div class="magnify"><a href="/wiki/File:Open_street_map_central_london.svg" title="File:Open street map central london.svg">class=notpageimage| </a></div>A map showing the location of the main campus of City, University of London, in central London</div></div></div>
<p>City has sites throughout London,<sup id="cite_ref-30" class="reference"><a href="#cite_note-30"><span class="cite-bracket">&#91;</span>30<span class="cite-bracket">&#93;</span></a></sup> with the main campus located at <a href="/wiki/Northampton_Square" title="Northampton Square">Northampton Square</a> in the <a href="/wiki/Finsbury" title="Finsbury">Finsbury</a> area of <a href="/wiki/London_Borough_of_Islington" title="London Borough of Islington">Islington</a>. The Rhind Building which houses the School of Arts and Social Sciences is directly west of Northampton Square. A few buildings of the main campus are located in nearby <a href="/wiki/Goswell_Road" title="Goswell Road">Goswell Road</a> in <a href="/wiki/Clerkenwell" title="Clerkenwell">Clerkenwell</a>.
</p><p>Other academic sites are:
</p>
<ul><li><a href="/wiki/City_Law_School" title="City Law School">The City Law School</a> (incorporating the former <i>Inns of Court School of Law</i>) on Sebastian Street, Islington.</li>
<li><a href="/wiki/Bayes_Business_School" title="Bayes Business School">Bayes Business School</a> in <a href="/wiki/St_Luke%27s,_London" title="St Luke&#39;s, London">St Luke's</a>, Islington, and at <a href="/wiki/200_Aldersgate" title="200 Aldersgate">200 Aldersgate</a> in <a href="/wiki/Smithfield,_London" title="Smithfield, London">Smithfield</a>, City of London</li>
<li>INTO City in <a href="/wiki/Spitalfields" title="Spitalfields">Spitalfields</a>, <a href="/wiki/London_Borough_of_Tower_Hamlets" title="London Borough of Tower Hamlets">Tower Hamlets</a></li></ul>
<div class="mw-heading mw-heading2"><h2 id="Organisation_and_administration">Organisation and administration</h2><span class="mw-editsection"><span class="mw-editsection-bracket">[</span><a href="/w/index.php?title=City,_University_of_London&amp;action=edit&amp;section=6" title="Edit section: Organisation and administration"><span>edit</span></a><span class="mw-editsection-bracket">]</span></span></div>
<figure class="mw-halign-right" typeof="mw:File/Thumb"><a href="/wiki/File:City_University_of_London_Northampton_Square_Clerkenwell_London_EC1V_0HB.jpg" class="mw-file-description"><img src="//upload.wikimedia.org/wikipedia/commons/thumb/2/29/City_University_of_London_Northampton_Square_Clerkenwell_London_EC1V_0HB.jpg/250px-City_University_of_London_Northampton_Square_Clerkenwell_London_EC1V_0HB.jpg" decoding="async" width="220" height="165" class="mw-file-element" srcset="//upload.wikimedia.org/wikipedia/commons/thumb/2/29/City_University_of_London_Northampton_Square_Clerkenwell_London_EC1V_0HB.jpg/330px-City_University_of_London_Northampton_Square_Clerkenwell_London_EC1V_0HB.jpg 1.5x, //upload.wikimedia.org/wikipedia/commons/thumb/2/29/City_University_of_London_Northampton_Square_Clerkenwell_London_EC1V_0HB.jpg/500px-City_University_of_London_Northampton_Square_Clerkenwell_London_EC1V_0HB.jpg 2x" data-file-width="4000" data-file-height="3000" /></a><figcaption>The main entrance of City St George's, University of London, in Northampton Square. The entrance was substantially remodelled in 2017 and opened by the Chancellor, <a href="/wiki/The_Princess_Royal" class="mw-redirect" title="The Princess Royal">The Princess Royal</a></figcaption></figure>
<p>The <a href="/wiki/Rector_(academia)" title="Rector (academia)">rector</a> of City St George's, University of London, is <i><a href="/wiki/Ex_officio" class="mw-redirect" title="Ex officio">ex officio</a></i> the <a href="/wiki/Lord_Mayor_of_the_City_of_London" class="mw-redirect" title="Lord Mayor of the City of London">Lord Mayor of the City of London</a>. The day-to-day running of the university is the responsibility of the <a href="/wiki/Vice-chancellor" title="Vice-chancellor">president</a>. The current president is Sir <a href="/wiki/Anthony_Finkelstein" title="Anthony Finkelstein">Anthony Finkelstein</a>.
</p>
<div class="mw-heading mw-heading3"><h3 id="Schools">Schools</h3><span class="mw-editsection"><span class="mw-editsection-bracket">[</span><a href="/w/index.php?title=City,_University_of_London&amp;action=edit&amp;section=7" title="Edit section: Schools"><span>edit</span></a><span class="mw-editsection-bracket">]</span></span></div>
<p>City St George's, University of London, is organised into six schools:
</p>
<ul><li><a href="/wiki/City_Law_School" title="City Law School">The City Law School</a>, incorporating <i>The Centre for Legal Studies</i> and the <a href="/wiki/Inns_of_Court_School_of_Law" class="mw-redirect" title="Inns of Court School of Law">Inns of Court School of Law</a></li>
<li><a href="/wiki/School_of_Health_Sciences,_City_University_London" class="mw-redirect" title="School of Health Sciences, City University London">School of Health &amp; Psychological Sciences</a>, incorporating <a href="/wiki/St_Bartholomew_School_of_Nursing_%26_Midwifery" class="mw-redirect" title="St Bartholomew School of Nursing &amp; Midwifery">St Bartholomew School of Nursing &amp; Midwifery</a></li>
<li>School of Communication &amp; Creativity, including the <a href="/wiki/City_University_Journalism_Department" class="mw-redirect" title="City University Journalism Department">Department of Journalism</a></li>
<li>School of Policy and Global Affairs</li>
<li>School of Science &amp; Technology</li>
<li><a href="/wiki/Bayes_Business_School" title="Bayes Business School">Bayes Business School</a> (formerly Cass Business School)</li></ul>
<div class="mw-heading mw-heading3"><h3 id="Finances">Finances</h3><span class="mw-editsection"><span class="mw-editsection-bracket">[</span><a href="/w/index.php?title=City,_University_of_London&amp;action=edit&amp;section=8" title="Edit section: Finances"><span>edit</span></a><span class="mw-editsection-bracket">]</span></span></div>
<p>In the financial year ended 31 July 2011, City had a total income (including share of joint ventures) of £178.6&#160;million (2008/09 – £174.4&#160;million) and total expenditure of £183.62&#160;million (2008/09 – £178.82&#160;million).<sup id="cite_ref-annrep_31-0" class="reference"><a href="#cite_note-annrep-31"><span class="cite-bracket">&#91;</span>31<span class="cite-bracket">&#93;</span></a></sup> Key sources of income included £39.58&#160;million from Funding Council grants (2008/09 – £39.52&#160;million), £116.91&#160;million from tuition fees and education contracts (2008/09 – £104.39&#160;million), £7.86&#160;million from research grants and contracts (2008/09 – £9.29&#160;million), £1.04 from endowment and investment income (2008/09 – £1.83&#160;million) and £15.05&#160;million from other income (2008/09 – £19.37&#160;million).<sup id="cite_ref-annrep_31-1" class="reference"><a href="#cite_note-annrep-31"><span class="cite-bracket">&#91;</span>31<span class="cite-bracket">&#93;</span></a></sup>
</p><p>During the 2010/11 financial year, City had a capital expenditure of £9.77&#160;million (2008/09 – £16.13&#160;million).<sup id="cite_ref-annrep_31-2" class="reference"><a href="#cite_note-annrep-31"><span class="cite-bracket">&#91;</span>31<span class="cite-bracket">&#93;</span></a></sup>
</p><p>At year end, City had reserves and endowments of £112.89&#160;million (2009/10 – £110.05&#160;million) and total net assets of £147.64&#160;million (2008/09 – £147.27&#160;million).<sup id="cite_ref-annrep_31-3" class="reference"><a href="#cite_note-annrep-31"><span class="cite-bracket">&#91;</span>31<span class="cite-bracket">&#93;</span></a></sup>
</p>
<div class="mw-heading mw-heading2"><h2 id="Academic_profile">Academic profile</h2><span class="mw-editsection"><span class="mw-editsection-bracket">[</span><a href="/w/index.php?title=City,_University_of_London&amp;action=edit&amp;section=9" title="Edit section: Academic profile"><span>edit</span></a><span class="mw-editsection-bracket">]</span></span></div>
<div class="mw-heading mw-heading3"><h3 id="Courses_and_rankings">Courses and rankings</h3><span class="mw-editsection"><span class="mw-editsection-bracket">[</span><a href="/w/index.php?title=City,_University_of_London&amp;action=edit&amp;section=10" title="Edit section: Courses and rankings"><span>edit</span></a><span class="mw-editsection-bracket">]</span></span></div>
<link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1289430074" /><table class="infobox"><caption class="infobox-title"><a href="/wiki/College_and_university_rankings" title="College and university rankings">Rankings</a></caption><tbody><tr><th colspan="2" class="infobox-header" style="background-color: #e0e0e0;"><a href="/wiki/Rankings_of_universities_in_the_United_Kingdom" title="Rankings of universities in the United Kingdom">National rankings</a></th></tr><tr><th scope="row" class="infobox-label"><i><a href="/wiki/Rankings_of_universities_in_the_United_Kingdom#The_Complete_University_Guide" title="Rankings of universities in the United Kingdom">Complete</a></i> (2025)<sup id="cite_ref-Complete_League_Table_32-0" class="reference"><a href="#cite_note-Complete_League_Table-32"><span class="cite-bracket">&#91;</span>32<span class="cite-bracket">&#93;</span></a></sup></th><td class="infobox-data">69</td></tr><tr><th scope="row" class="infobox-label"><i><a href="/wiki/Rankings_of_universities_in_the_United_Kingdom#The_Guardian" title="Rankings of universities in the United Kingdom">Guardian</a></i> (2025)<sup id="cite_ref-The_Guardian_University_Guide_33-0" class="reference"><a href="#cite_note-The_Guardian_University_Guide-33"><span class="cite-bracket">&#91;</span>33<span class="cite-bracket">&#93;</span></a></sup></th><td class="infobox-data">38</td></tr><tr><th scope="row" class="infobox-label"><i><a href="/wiki/Rankings_of_universities_in_the_United_Kingdom#The_Times/The_Sunday_Times" title="Rankings of universities in the United Kingdom">Times / Sunday Times</a></i> (2025)<sup id="cite_ref-The_Times_and_Sunday_Times_University_Guide_34-0" class="reference"><a href="#cite_note-The_Times_and_Sunday_Times_University_Guide-34"><span class="cite-bracket">&#91;</span>34<span class="cite-bracket">&#93;</span></a></sup></th><td class="infobox-data">49</td></tr><tr><th colspan="2" class="infobox-header" style="background-color: #e0e0e0;">Global rankings</th></tr><tr><th scope="row" class="infobox-label"><i><a href="/wiki/Academic_Ranking_of_World_Universities" title="Academic Ranking of World Universities">ARWU</a></i> (2024)<sup id="cite_ref-Academic_Ranking_of_World_Universities_35-0" class="reference"><a href="#cite_note-Academic_Ranking_of_World_Universities-35"><span class="cite-bracket">&#91;</span>35<span class="cite-bracket">&#93;</span></a></sup></th><td class="infobox-data">901–1000</td></tr><tr><th scope="row" class="infobox-label"><i><a href="/wiki/QS_World_University_Rankings" title="QS World University Rankings">QS</a></i> (2025)<sup id="cite_ref-QS_World_University_Rankings_36-0" class="reference"><a href="#cite_note-QS_World_University_Rankings-36"><span class="cite-bracket">&#91;</span>36<span class="cite-bracket">&#93;</span></a></sup></th><td class="infobox-data">352</td></tr><tr><th scope="row" class="infobox-label"><i><a href="/wiki/Times_Higher_Education_World_University_Rankings" title="Times Higher Education World University Rankings">THE</a></i> (2025)<sup id="cite_ref-THE_World_University_Rankings_37-0" class="reference"><a href="#cite_note-THE_World_University_Rankings-37"><span class="cite-bracket">&#91;</span>37<span class="cite-bracket">&#93;</span></a></sup></th><td class="infobox-data">351–400</td></tr></tbody></table>
<p>City St George's, University of London, offers Bachelor's, Master's, and <a href="/wiki/Doctorate" title="Doctorate">Doctoral</a> <a href="/wiki/Academic_degree" title="Academic degree">degrees</a> as well as certificates and diplomas at both undergraduate and postgraduate level. More than two-thirds of City's programmes are recognised by the appropriate <a href="/wiki/Professional_bodies" class="mw-redirect" title="Professional bodies">professional bodies</a> such as the <a href="/wiki/British_Computer_Society" title="British Computer Society">BCS</a>, <a href="/wiki/British_Psychological_Society" title="British Psychological Society">BPS</a>, <a href="/wiki/Chartered_Institute_of_Library_and_Information_Professionals" title="Chartered Institute of Library and Information Professionals">CILIP</a>, <a href="/wiki/Institution_of_Civil_Engineers" title="Institution of Civil Engineers">ICE</a>, <a href="/wiki/Royal_Institution_of_Chartered_Surveyors" title="Royal Institution of Chartered Surveyors">RICS</a>, <a href="/wiki/Health_Professions_Council" class="mw-redirect" title="Health Professions Council">HPC</a> etc. in recognition of the high standards of relevance to the professions. The university also has an online careers network where over 2,000 former students offer practical help to current students.<sup id="cite_ref-38" class="reference"><a href="#cite_note-38"><span class="cite-bracket">&#91;</span>38<span class="cite-bracket">&#93;</span></a></sup>
</p><p>The <a href="/wiki/City_Law_School" title="City Law School">City Law School</a> offers courses for undergraduates, postgraduates, master graduates and professional courses leading to qualification as a solicitor or barrister, as well as continuing professional development. Its Legal Practice Course has the highest quality rating from the <a href="/wiki/Solicitors_Regulation_Authority" title="Solicitors Regulation Authority">Solicitors Regulation Authority</a>.<sup id="cite_ref-39" class="reference"><a href="#cite_note-39"><span class="cite-bracket">&#91;</span>39<span class="cite-bracket">&#93;</span></a></sup>
</p><p>The <a href="/wiki/School_of_Community_and_Health_Sciences#Department_of_Radiography" class="mw-redirect" title="School of Community and Health Sciences">Department of Radiography</a> (part of the <a href="/wiki/School_of_Community_and_Health_Sciences" class="mw-redirect" title="School of Community and Health Sciences">School of Community and Health Sciences</a>) offers two radiography degrees, the <a href="/wiki/School_of_Community_and_Health_Sciences#Diagnostic_Radiography" class="mw-redirect" title="School of Community and Health Sciences">BSc (Hons) Radiography (Diagnostic Imaging)</a> and BSc (Hons) Radiography (Radiotherapy and Oncology), both of which are recognised by the <a href="/wiki/Health_Professions_Council" class="mw-redirect" title="Health Professions Council">Health Professions Council</a> (HPC).
</p>
<div class="mw-heading mw-heading3"><h3 id="Partnerships_and_collaborations">Partnerships and collaborations</h3><span class="mw-editsection"><span class="mw-editsection-bracket">[</span><a href="/w/index.php?title=City,_University_of_London&amp;action=edit&amp;section=11" title="Edit section: Partnerships and collaborations"><span>edit</span></a><span class="mw-editsection-bracket">]</span></span></div>
<div class="mw-heading mw-heading4"><h4 id="CETL">CETL</h4><span class="mw-editsection"><span class="mw-editsection-bracket">[</span><a href="/w/index.php?title=City,_University_of_London&amp;action=edit&amp;section=12" title="Edit section: CETL"><span>edit</span></a><span class="mw-editsection-bracket">]</span></span></div>
<p><a href="/wiki/Queen_Mary,_University_of_London" class="mw-redirect" title="Queen Mary, University of London">Queen Mary, University of London</a>, and City St George's, University of London, were jointly awarded Centre for Excellence in Teaching and Learning (CETL) status by the <a href="/wiki/Higher_Education_Funding_Council_for_England" title="Higher Education Funding Council for England">Higher Education Funding Council for England</a> (HEFCE) in recognition of their work in skills training for 3,000 students across six healthcare professions.<sup id="cite_ref-40" class="reference"><a href="#cite_note-40"><span class="cite-bracket">&#91;</span>40<span class="cite-bracket">&#93;</span></a></sup>
</p>
<div class="mw-heading mw-heading4"><h4 id="City_of_London">City of London</h4><span class="mw-editsection"><span class="mw-editsection-bracket">[</span><a href="/w/index.php?title=City,_University_of_London&amp;action=edit&amp;section=13" title="Edit section: City of London"><span>edit</span></a><span class="mw-editsection-bracket">]</span></span></div>
<p>City St George's, University of London, has links with businesses in the <a href="/wiki/City_of_London" title="City of London">City of London</a>.<sup id="cite_ref-:0_41-0" class="reference"><a href="#cite_note-:0-41"><span class="cite-bracket">&#91;</span>41<span class="cite-bracket">&#93;</span></a></sup> City has also joined forces with other universities such as <a href="/wiki/Queen_Mary,_University_of_London" class="mw-redirect" title="Queen Mary, University of London">Queen Mary</a> and the <a href="/wiki/Institute_of_Education" class="mw-redirect" title="Institute of Education">Institute of Education</a> (both part of the <a href="/wiki/University_of_London" title="University of London">University of London</a>) with which it jointly delivers several leading degree programmes.
</p>
<div class="mw-heading mw-heading4"><h4 id="LCACE">LCACE</h4><span class="mw-editsection"><span class="mw-editsection-bracket">[</span><a href="/w/index.php?title=City,_University_of_London&amp;action=edit&amp;section=14" title="Edit section: LCACE"><span>edit</span></a><span class="mw-editsection-bracket">]</span></span></div>
<p><a href="/wiki/LCACE" class="mw-redirect" title="LCACE">London Centre for Arts and Cultural Exchange</a> is a consortium of nine universities. It was established in 2004 to foster collaboration and to promote and support the exchange of knowledge between the consortium's partners and London's arts and cultural sectors. The nine institutions involved are: <a href="/wiki/University_of_the_Arts_London" title="University of the Arts London">University of the Arts London</a>; <a href="/wiki/Birkbeck,_University_of_London" title="Birkbeck, University of London">Birkbeck, University of London</a>; City St George's, University of London; <a href="/wiki/The_Courtauld_Institute_of_Art" class="mw-redirect" title="The Courtauld Institute of Art">The Courtauld Institute of Art</a>; <a href="/wiki/Goldsmiths,_University_of_London" title="Goldsmiths, University of London">Goldsmiths, University of London</a>; <a href="/wiki/Guildhall_School_of_Music_%26_Drama" class="mw-redirect" title="Guildhall School of Music &amp; Drama">Guildhall School of Music &amp; Drama</a>; <a href="/wiki/King%27s_College_London" title="King&#39;s College London">King's College London</a>; <a href="/wiki/Queen_Mary,_University_of_London" class="mw-redirect" title="Queen Mary, University of London">Queen Mary, University of London</a>, and <a href="/wiki/Royal_Holloway,_University_of_London" title="Royal Holloway, University of London">Royal Holloway, University of London</a>.
</p>
<div class="mw-heading mw-heading4"><h4 id="WC2_University_Network">WC2 University Network</h4><span class="mw-editsection"><span class="mw-editsection-bracket">[</span><a href="/w/index.php?title=City,_University_of_London&amp;action=edit&amp;section=15" title="Edit section: WC2 University Network"><span>edit</span></a><span class="mw-editsection-bracket">]</span></span></div>
<p>City is a founding member of the WC2 University Network, a network of universities developed with the goal of bringing together leading universities located in the heart of major world cities in order to address cultural, environmental and political issues of common interest to world cities and their universities.<sup id="cite_ref-:1_13-1" class="reference"><a href="#cite_note-:1-13"><span class="cite-bracket">&#91;</span>13<span class="cite-bracket">&#93;</span></a></sup> In addition to City St George's, University of London, the founding members of WC2 members are: <a href="/wiki/City_University_of_New_York" title="City University of New York">City University of New York</a>, <a href="/wiki/Technische_Universit%C3%A4t_Berlin" title="Technische Universität Berlin">Technische Universität Berlin</a>, <a href="/wiki/Universidade_de_S%C3%A3o_Paulo" class="mw-redirect" title="Universidade de São Paulo">Universidade de São Paulo</a>, <a href="/wiki/Hong_Kong_Polytechnic_University" title="Hong Kong Polytechnic University">Hong Kong Polytechnic University</a>, <a href="/wiki/Universidad_Autonoma_Metropolitana" class="mw-redirect" title="Universidad Autonoma Metropolitana">Universidad Autonoma Metropolitana</a>, <a href="/wiki/Saint_Petersburg_State_Polytechnical_University" class="mw-redirect" title="Saint Petersburg State Polytechnical University">Saint Petersburg State Polytechnical University</a>, <a href="/wiki/Politecnico_di_Milano" class="mw-redirect" title="Politecnico di Milano">Politecnico di Milano</a>, <a href="/wiki/University_of_Delhi" class="mw-redirect" title="University of Delhi">University of Delhi</a>, <a href="/wiki/Northeastern_University" title="Northeastern University">Northeastern University</a> Boston and <a href="/wiki/Tongji_University" title="Tongji University">Tongji University</a>.
</p>
<div class="mw-heading mw-heading4"><h4 id="Erasmus_Mundus_MULTI">Erasmus Mundus MULTI</h4><span class="mw-editsection"><span class="mw-editsection-bracket">[</span><a href="/w/index.php?title=City,_University_of_London&amp;action=edit&amp;section=16" title="Edit section: Erasmus Mundus MULTI"><span>edit</span></a><span class="mw-editsection-bracket">]</span></span></div>
<p>City was selected as the sole British university to take part in the selective Erasmus Mundus MULTI programme, funded by the <a href="/wiki/European_Commission" title="European Commission">European Commission</a> to promote scientific exchange between Europe and the industrialised countries of South-East Asia. It is the first Erasmus program to involve universities outside of Europe. In addition to City, the partner universities are: <a href="/wiki/Aix-Marseille_University" title="Aix-Marseille University">Aix-Marseille University</a> (France), <a href="/wiki/Univerzita_Karlova_v_Praze" class="mw-redirect" title="Univerzita Karlova v Praze">Univerzita Karlova v Praze</a> (Czech Republic), <a href="/wiki/Freie_Universit%C3%A4t_Berlin" class="mw-redirect" title="Freie Universität Berlin">Freie Universität Berlin</a> (Germany), <a href="/wiki/Universit%C3%A4t_des_Saarlandes" class="mw-redirect" title="Universität des Saarlandes">Universität des Saarlandes</a> (Germany), <a href="/wiki/Universit%C3%A0_di_Pisa" class="mw-redirect" title="Università di Pisa">Università di Pisa</a> (Italy), <a href="/wiki/Universidad_de_Sevilla" class="mw-redirect" title="Universidad de Sevilla">Universidad de Sevilla</a> (Spain), <a href="/wiki/The_Hong_Kong_Polytechnic_University" class="mw-redirect" title="The Hong Kong Polytechnic University">The Hong Kong Polytechnic University</a> (Hong Kong, SAR China), <a href="/wiki/Universiti_Brunei_Darussalam" title="Universiti Brunei Darussalam">Universiti Brunei Darussalam</a> (Brunei), <a href="/wiki/University_of_Macau" title="University of Macau">University of Macau</a> (Macau, SAR China), <a href="/wiki/Nanyang_Technological_University" title="Nanyang Technological University">Nanyang Technological University</a> (Singapore), and <a href="/wiki/National_Taiwan_University" title="National Taiwan University">National Taiwan University</a> (Taiwan).
</p>
<div class="mw-heading mw-heading4"><h4 id="UCL_Partners">UCL Partners</h4><span class="mw-editsection"><span class="mw-editsection-bracket">[</span><a href="/w/index.php?title=City,_University_of_London&amp;action=edit&amp;section=17" title="Edit section: UCL Partners"><span>edit</span></a><span class="mw-editsection-bracket">]</span></span></div>
<p>City has joined the executive group of <a href="/wiki/UCL_Partners" class="mw-redirect" title="UCL Partners">UCL Partners</a>, one of five accredited academic health science groups in the UK. City was invited to join the partnership in recognition of its expertise in nursing, allied health, health services research and evaluation and health management.<sup id="cite_ref-42" class="reference"><a href="#cite_note-42"><span class="cite-bracket">&#91;</span>42<span class="cite-bracket">&#93;</span></a></sup>
</p>
<div class="mw-heading mw-heading4"><h4 id="City_Research_Online">City Research Online</h4><span class="mw-editsection"><span class="mw-editsection-bracket">[</span><a href="/w/index.php?title=City,_University_of_London&amp;action=edit&amp;section=18" title="Edit section: City Research Online"><span>edit</span></a><span class="mw-editsection-bracket">]</span></span></div>
<p><i>City Research Online</i> provides <a href="/wiki/Open_access" title="Open access">open access</a> to, and reliable information about, research produced by City staff and research students, as permitted by publishers and copyright law, of content and <a href="/wiki/Metadata" title="Metadata">metadata</a>.<sup id="cite_ref-openaccess.city.ac.uk/information_43-0" class="reference"><a href="#cite_note-openaccess.city.ac.uk/information-43"><span class="cite-bracket">&#91;</span>43<span class="cite-bracket">&#93;</span></a></sup>
</p><p>These<sup id="cite_ref-openaccess.city.ac.uk/information_43-1" class="reference"><a href="#cite_note-openaccess.city.ac.uk/information-43"><span class="cite-bracket">&#91;</span>43<span class="cite-bracket">&#93;</span></a></sup> include:
</p>
<ul><li><a href="/wiki/Academic_publishing#Scholarly_paper" title="Academic publishing">Articles</a> (submitted, accepted and published versions)</li>
<li><a href="/wiki/Working_paper" title="Working paper">Working papers</a></li>
<li><a href="/wiki/Book" title="Book">Books</a></li>
<li><a href="/wiki/Book_chapter" class="mw-redirect" title="Book chapter">Book chapters</a></li>
<li><a href="/wiki/Conference_paper" class="mw-redirect" title="Conference paper">Conference papers</a></li>
<li><a href="/wiki/Multimedia" title="Multimedia">Multimedia</a></li>
<li><a href="/wiki/Doctorate" title="Doctorate">Doctoral</a> <a href="/wiki/Thesis" title="Thesis">theses</a></li></ul>
<div class="mw-heading mw-heading2"><h2 id="Student_life">Student life</h2><span class="mw-editsection"><span class="mw-editsection-bracket">[</span><a href="/w/index.php?title=City,_University_of_London&amp;action=edit&amp;section=19" title="Edit section: Student life"><span>edit</span></a><span class="mw-editsection-bracket">]</span></span></div>
<div class="mw-heading mw-heading3"><h3 id="Students'_Union"><span id="Students.27_Union"></span>Students' Union</h3><span class="mw-editsection"><span class="mw-editsection-bracket">[</span><a href="/w/index.php?title=City,_University_of_London&amp;action=edit&amp;section=20" title="Edit section: Students&#039; Union"><span>edit</span></a><span class="mw-editsection-bracket">]</span></span></div>
<p>The City Students' Union is run primarily by students through four elected sabbatical officers, the chief executive and an elected assembly (composed of current students), with oversight by a trustee board. The Students' Union provides support, representation, facilities, services, entertainment and activities for its members. It is run for students, by students.<sup id="cite_ref-44" class="reference"><a href="#cite_note-44"><span class="cite-bracket">&#91;</span>44<span class="cite-bracket">&#93;</span></a></sup>
</p><p>The Students' Union manages most aspects relating to students' societies, such as booking spaces for events on campus, holding funds and distributing grants, and providing training to their committees.  
</p>
<div class="mw-heading mw-heading3"><h3 id="Student_media">Student media</h3><span class="mw-editsection"><span class="mw-editsection-bracket">[</span><a href="/w/index.php?title=City,_University_of_London&amp;action=edit&amp;section=21" title="Edit section: Student media"><span>edit</span></a><span class="mw-editsection-bracket">]</span></span></div>
<p>City currently has two student-run media outlets, including Carrot Radio, which was co-founded by journalism postgraduates Jordan Gass-Pooré and Winston Lo in the autumn of 2018.<sup id="cite_ref-45" class="reference"><a href="#cite_note-45"><span class="cite-bracket">&#91;</span>45<span class="cite-bracket">&#93;</span></a></sup> Carrot Radio currently records weekday podcasts. The second is the student-led online magazine, <i>Carrot Magazine</i>. They released their first print magazine in December 2017.
</p>
<div class="mw-heading mw-heading3"><h3 id="Other">Other</h3><span class="mw-editsection"><span class="mw-editsection-bracket">[</span><a href="/w/index.php?title=City,_University_of_London&amp;action=edit&amp;section=22" title="Edit section: Other"><span>edit</span></a><span class="mw-editsection-bracket">]</span></span></div>
<p>For a number of years, City students have taken part in the annual <a href="/wiki/Lord_Mayor%27s_Show" title="Lord Mayor&#39;s Show">Lord Mayor's Show</a>, representing the university in one of the country's largest and liveliest parades.
</p>
<div class="mw-heading mw-heading2"><h2 id="Sustainability_ranking">Sustainability ranking</h2><span class="mw-editsection"><span class="mw-editsection-bracket">[</span><a href="/w/index.php?title=City,_University_of_London&amp;action=edit&amp;section=23" title="Edit section: Sustainability ranking"><span>edit</span></a><span class="mw-editsection-bracket">]</span></span></div>
<p>City ranked joint 5th out of the 168 universities surveyed in the 2019 <a href="/wiki/People_%26_Planet" title="People &amp; Planet">People &amp; Planet</a> league table of the most sustainable UK universities<sup id="cite_ref-46" class="reference"><a href="#cite_note-46"><span class="cite-bracket">&#91;</span>46<span class="cite-bracket">&#93;</span></a></sup> having climbed from 7th place in the 2016 league. In both the 2016 and 2019 rankings, it was the highest ranking University of London institution, and one of only four London institutions in the top twenty.
</p><p>The league table's <i>Fossil Free Scorecard</i> report, drawn from <a href="/wiki/Freedom_of_information_in_the_United_Kingdom" title="Freedom of information in the United Kingdom">Freedom of Information</a> requests, found that £800,000 (6.4%) of City's £12.5m endowment was invested in <a href="/wiki/Fossil_fuels" class="mw-redirect" title="Fossil fuels">fossil fuels</a>, and that the institution had not made a public commitment to <a href="/wiki/Fossil_fuel_divestment" title="Fossil fuel divestment">fossil fuel divestment</a>. It also noted nearly £1m of research funding into <a href="/wiki/Renewable_energy" title="Renewable energy">renewables</a> since 2001 with just £64k of total funding from fossil fuel companies; and no <a href="/wiki/Honorary_degree" title="Honorary degree">honorary degrees</a> or board positions held by fossil fuel executives.<sup id="cite_ref-47" class="reference"><a href="#cite_note-47"><span class="cite-bracket">&#91;</span>47<span class="cite-bracket">&#93;</span></a></sup>
</p><p>City announced on 4 July 2023 that it was divesting its investments from fossil fuel producers.<sup id="cite_ref-48" class="reference"><a href="#cite_note-48"><span class="cite-bracket">&#91;</span>48<span class="cite-bracket">&#93;</span></a></sup>
</p>
<div class="mw-heading mw-heading2"><h2 id="Notable_people">Notable people</h2><span class="mw-editsection"><span class="mw-editsection-bracket">[</span><a href="/w/index.php?title=City,_University_of_London&amp;action=edit&amp;section=24" title="Edit section: Notable people"><span>edit</span></a><span class="mw-editsection-bracket">]</span></span></div>
<div class="mw-heading mw-heading3"><h3 id="Notable_alumni">Notable alumni</h3><span class="mw-editsection"><span class="mw-editsection-bracket">[</span><a href="/w/index.php?title=City,_University_of_London&amp;action=edit&amp;section=25" title="Edit section: Notable alumni"><span>edit</span></a><span class="mw-editsection-bracket">]</span></span></div>
<div class="mw-heading mw-heading4"><h4 id="Government,_politics_and_society"><span id="Government.2C_politics_and_society"></span>Government, politics and society</h4><span class="mw-editsection"><span class="mw-editsection-bracket">[</span><a href="/w/index.php?title=City,_University_of_London&amp;action=edit&amp;section=26" title="Edit section: Government, politics and society"><span>edit</span></a><span class="mw-editsection-bracket">]</span></span></div>
<figure class="mw-default-size" typeof="mw:File/Thumb"><a href="/wiki/File:Clement_Attlee.jpg" class="mw-file-description"><img src="//upload.wikimedia.org/wikipedia/commons/thumb/1/1e/Clement_Attlee.jpg/250px-Clement_Attlee.jpg" decoding="async" width="250" height="357" class="mw-file-element" srcset="//upload.wikimedia.org/wikipedia/commons/thumb/1/1e/Clement_Attlee.jpg/500px-Clement_Attlee.jpg 1.5x" data-file-width="1712" data-file-height="2445" /></a><figcaption><a href="/wiki/Clement_Attlee" title="Clement Attlee">Clement Attlee</a></figcaption></figure>
<ul><li><a href="/wiki/Muhammad_Ali_Jinnah" title="Muhammad Ali Jinnah">Muhammad Ali Jinnah</a> – founder of Pakistan, first <a href="/wiki/Governor-General_of_Pakistan" title="Governor-General of Pakistan">Governor-General of Pakistan</a> graduated from the Inns of Court School of Law (now part of <a href="/wiki/City_Law_School" title="City Law School">The City Law School</a>)</li>
<li><a href="/wiki/Clement_Attlee" title="Clement Attlee">Clement Attlee</a> – <a href="/wiki/Labour_Party_(UK)" title="Labour Party (UK)">Labour</a> <a href="/wiki/Prime_Minister_of_the_United_Kingdom" title="Prime Minister of the United Kingdom">Prime Minister of the United Kingdom</a> from 1945 to 1951</li>
<li><a href="/wiki/H._H._Asquith" title="H. H. Asquith">H. H. Asquith</a> – <a href="/wiki/Liberal_Party_(UK)" title="Liberal Party (UK)">Liberal</a> Prime Minister of the United Kingdom from 1908 to 1916</li>
<li><a href="/wiki/Sir_Tony_Blair" class="mw-redirect" title="Sir Tony Blair">Sir Tony Blair</a> – Labour Party Prime Minister of the United Kingdom from 1997 to 2007, graduated from the Inns of Court School of Law (now part of <a href="/wiki/City_Law_School" title="City Law School">The City Law School</a>)</li>
<li><a href="/wiki/Christos_Staikouras" title="Christos Staikouras">Christos Staikouras</a> – Finance Minister of Greece from 2019 to present</li>
<li><a href="/wiki/Roderic_Bowen" title="Roderic Bowen">Roderic Bowen</a> – Welsh Liberal Party politician</li>
<li><a href="/wiki/Robert_Chote" title="Robert Chote">Robert Chote</a> – chief of the <a href="/wiki/Office_for_Budget_Responsibility" title="Office for Budget Responsibility">Office for Budget Responsibility</a>; former director of <a href="/wiki/Institute_for_Fiscal_Studies" title="Institute for Fiscal Studies">Institute for Fiscal Studies</a></li>
<li><a href="/wiki/Ali_Dizaei" title="Ali Dizaei">Ali Dizaei</a> – former police commander</li>
<li><a href="/wiki/Sir_James_Dutton" class="mw-redirect" title="Sir James Dutton">Sir James Dutton</a> – Royal Marine general and former deputy commander of the <a href="/wiki/International_Security_Assistance_Force" title="International Security Assistance Force">International Security Assistance Force</a></li>
<li><a href="/wiki/Chlo%C3%AB_Fox" title="Chloë Fox">Chloë Fox</a> – Australian politician, former <a href="/wiki/Australian_Labor_Party" title="Australian Labor Party">Labor</a> MP for the South Australian <a href="/wiki/Electoral_district_of_Bright" title="Electoral district of Bright">electoral district of Bright</a></li>
<li><a href="/wiki/Mahatma_Gandhi" title="Mahatma Gandhi">Mahatma Gandhi</a> – Leader of the <a href="/wiki/Indian_Independence_Movement" class="mw-redirect" title="Indian Independence Movement">Indian Independence Movement</a>, graduated in 1891 from the Inns of Court School of Law (now part of <a href="/wiki/City_Law_School" title="City Law School">The City Law School</a>)</li>
<li><a href="/wiki/Syed_Sharifuddin_Pirzada" class="mw-redirect" title="Syed Sharifuddin Pirzada">Syed Sharifuddin Pirzada</a> – Noted Pakistani lawyer &amp; Politician. Also served as 5th secretary general of <a href="/wiki/Organisation_of_Islamic_Cooperation" title="Organisation of Islamic Cooperation">Organisation of Islamic Cooperation</a>.</li>
<li><a href="/wiki/James_Hart_(police_commissioner)" class="mw-redirect" title="James Hart (police commissioner)">James Hart</a> – Commissioner of the <a href="/wiki/City_of_London_Police" title="City of London Police">City of London Police</a></li>
<li><a href="/wiki/David_Heath_(politician)" title="David Heath (politician)">David Heath</a> – Politician and Liberal Democrat Member of Parliament for <a href="/wiki/Somerton_and_Frome_(UK_Parliament_constituency)" title="Somerton and Frome (UK Parliament constituency)">Somerton and Frome</a></li>
<li><a href="/wiki/Syed_Kamall" title="Syed Kamall">Syed Kamall</a> – <a href="/wiki/Conservative_Party_(UK)" title="Conservative Party (UK)">Conservative Party</a> politician and <a href="/wiki/Member_of_the_European_Parliament" title="Member of the European Parliament">Member of the European Parliament</a> for the <a href="/wiki/London_(European_Parliament_constituency)" title="London (European Parliament constituency)">London European Parliament constituency</a></li>
<li><a href="/wiki/David_Lammy" title="David Lammy">David Lammy</a> – Labour MP for Tottenham</li>
<li><a href="/wiki/Sandro_Marcos" title="Sandro Marcos">Sandro Marcos</a> – Member of the Philippine House of Representatives, eldest son of <a href="/wiki/Bongbong_Marcos" title="Bongbong Marcos">Bongbong Marcos</a>, <a href="/wiki/President_of_the_Philippines" title="President of the Philippines">President of the Republic of the Philippines</a></li>
<li><a href="/wiki/Barbara_Mensah" title="Barbara Mensah">Barbara Mensah</a> – Judge<sup id="cite_ref-bbc_49-0" class="reference"><a href="#cite_note-bbc-49"><span class="cite-bracket">&#91;</span>49<span class="cite-bracket">&#93;</span></a></sup></li>
<li><a href="/wiki/Liu_Mingkang" title="Liu Mingkang">Liu Mingkang</a> – Chinese Politician and Businessman, current Chairman of the <a href="/wiki/China_Banking_Regulatory_Commission" title="China Banking Regulatory Commission">China Banking Regulatory Commission</a>, former Vice-Governor of the <a href="/wiki/China_Development_Bank" title="China Development Bank">China Development Bank</a></li>
<li><a href="/wiki/Jawaharlal_Nehru" title="Jawaharlal Nehru">Jawaharlal Nehru</a> – First <a href="/wiki/Prime_Minister_of_India" title="Prime Minister of India">Prime Minister</a> of the Republic of India</li>
<li><a href="/wiki/Houda_Nonoo" title="Houda Nonoo">Houda Nonoo</a> – Bahraini Ambassador to the United States</li>
<li><a href="/wiki/Patrick_O%27Flynn" title="Patrick O&#39;Flynn">Patrick O'Flynn</a> – <a href="/wiki/UK_Independence_Party" title="UK Independence Party">UK Independence Party</a> MEP</li>
<li><a href="/wiki/Stav_Shaffir" title="Stav Shaffir">Stav Shaffir</a> – Youngest member of the Israeli <a href="/wiki/Knesset" title="Knesset">Knesset</a>, leader of the <a href="/wiki/Social_justice" title="Social justice">social justice</a> movement</li>
<li><a href="/wiki/Aris_Spiliotopoulos" title="Aris Spiliotopoulos">Aris Spiliotopoulos</a> – Minister of Greek Tourism</li>
<li><a href="/wiki/Margaret_Thatcher" title="Margaret Thatcher">Margaret Thatcher</a> – Conservative Party Prime Minister of the United Kingdom from 1979 to 1990, graduated from the Inns of Court School of Law (now part of <a href="/wiki/City_Law_School" title="City Law School">The City Law School</a>)</li>
<li><a href="/wiki/Ivy_Williams" title="Ivy Williams">Ivy Williams</a> – First woman to be called to the <a href="/wiki/English_bar" class="mw-redirect" title="English bar">English bar</a></li></ul>
<div class="mw-heading mw-heading4"><h4 id="Arts,_science_and_academia"><span id="Arts.2C_science_and_academia"></span>Arts, science and academia</h4><span class="mw-editsection"><span class="mw-editsection-bracket">[</span><a href="/w/index.php?title=City,_University_of_London&amp;action=edit&amp;section=27" title="Edit section: Arts, science and academia"><span>edit</span></a><span class="mw-editsection-bracket">]</span></span></div>
<figure class="mw-default-size mw-halign-right" typeof="mw:File/Thumb"><a href="/wiki/File:Iqbal.jpg" class="mw-file-description"><img src="//upload.wikimedia.org/wikipedia/commons/thumb/8/8a/Iqbal.jpg/190px-Iqbal.jpg" decoding="async" width="190" height="251" class="mw-file-element" srcset="//upload.wikimedia.org/wikipedia/commons/8/8a/Iqbal.jpg 1.5x" data-file-width="227" data-file-height="300" /></a><figcaption><a href="/wiki/Muhammad_Iqbal" title="Muhammad Iqbal">Muhammad Iqbal</a></figcaption></figure>
<ul><li><a href="/wiki/Darshana_Rajendran" title="Darshana Rajendran">Darshana Rajendran</a> - Indian Actress</li>
<li><a href="/wiki/L._Bruce_Archer" title="L. Bruce Archer">L. Bruce Archer</a> – British mechanical engineer and Professor of Design Research at the <a href="/wiki/Royal_College_of_Art" title="Royal College of Art">Royal College of Art</a></li>
<li><a href="/wiki/H%C3%BCseyin_%C5%9Eehito%C4%9Flu" title="Hüseyin Şehitoğlu">Hüseyin Şehitoğlu</a> –  Turkish Professor of Mechanical Science and Engineering, former department head, University of Illinois</li>
<li><a href="/wiki/Susan_Bickley" title="Susan Bickley">Susan Bickley</a> – Mezzo-soprano in opera and classical music</li>
<li><a href="/wiki/George_Daniels_(watchmaker)" title="George Daniels (watchmaker)">George Daniels</a> – <a href="/wiki/Horologist" class="mw-redirect" title="Horologist">Horologist</a> and inventor of the <a href="/wiki/Coaxial_escapement" title="Coaxial escapement">coaxial escapement</a></li>
<li><a href="/wiki/Jerry_Fishenden" title="Jerry Fishenden">Jerry Fishenden</a> – Technologist, former Microsoft National Technology Officer for the UK<sup id="cite_ref-50" class="reference"><a href="#cite_note-50"><span class="cite-bracket">&#91;</span>50<span class="cite-bracket">&#93;</span></a></sup></li>
<li><a href="/wiki/Julia_Gomelskaya" title="Julia Gomelskaya">Julia Gomelskaya</a> – Ukrainian contemporary music composer, professor of Odesa State Music Academy in Ukraine</li>
<li><a href="/wiki/Norman_Gowar" title="Norman Gowar">Norman Gowar</a> – Professor of Mathematics at the <a href="/wiki/Open_University" title="Open University">Open University</a> and Principal of <a href="/wiki/Royal_Holloway_College" class="mw-redirect" title="Royal Holloway College">Royal Holloway College</a>, University of London</li>
<li><a href="/wiki/Clare_Hammond" title="Clare Hammond">Clare Hammond</a> – Concert pianist</li>
<li><a href="/wiki/David_Hirsh" title="David Hirsh">David Hirsh</a> – Academic and sociologist</li>
<li><a href="/wiki/Muhammad_Iqbal" title="Muhammad Iqbal">Muhammad Iqbal</a> – Muslim poet, philosopher and politician, born in present-day Pakistan, graduated from the Inns of Court School of Law and University of Cambridge</li>
<li><a href="/wiki/John_Hodge_(engineer)" title="John Hodge (engineer)">John Hodge</a> – Aeronautical Engineer who played a key role in <a href="/wiki/NASA" title="NASA">NASA</a> and America's space race.<sup id="cite_ref-51" class="reference"><a href="#cite_note-51"><span class="cite-bracket">&#91;</span>51<span class="cite-bracket">&#93;</span></a></sup></li>
<li><a href="/wiki/John_Loder_(sound_engineer)" title="John Loder (sound engineer)">John Loder</a> – Sound engineer, record producer and founder of <a href="/wiki/Southern_Studios" title="Southern Studios">Southern Studios</a>, as well as a former member of <a href="/wiki/EXIT_(performance_art_group)" title="EXIT (performance art group)">EXIT</a></li>
<li><a href="/wiki/Sharon_Maguire" title="Sharon Maguire">Sharon Maguire</a> – Director of <i><a href="/wiki/Bridget_Jones%27s_Diary_(film)" class="mw-redirect" title="Bridget Jones&#39;s Diary (film)">Bridget Jones's Diary</a></i></li>
<li><a href="/wiki/Rhodri_Marsden" title="Rhodri Marsden">Rhodri Marsden</a> – Journalist, musician and blogger; columnist for <i><a href="/wiki/The_Independent" title="The Independent">The Independent</a></i></li>
<li><a href="/wiki/Robin_Milner" title="Robin Milner">Robin Milner</a> – Computer scientist and recipient of the 1991 <a href="/wiki/Association_for_Computing_Machinery" title="Association for Computing Machinery">ACM</a> <a href="/wiki/Turing_Award" title="Turing Award">Turing Award</a></li>
<li><a href="/wiki/Bernard_Miles" title="Bernard Miles">Bernard Miles</a> – Actor and founder of the <a href="/wiki/Mermaid_Theatre" title="Mermaid Theatre">Mermaid Theatre</a>.</li>
<li><a href="/wiki/John_Palmer_(20th-century_composer)" class="mw-redirect" title="John Palmer (20th-century composer)">John Palmer</a> – Instrumental and electroacoustic music composer</li>
<li><a href="/wiki/Ziauddin_Sardar" title="Ziauddin Sardar">Ziauddin Sardar</a> – Academic and scholar of Islamic issues, Commissioner of the <a href="/wiki/Equality_and_Human_Rights_Commission" title="Equality and Human Rights Commission">Equality and Human Rights Commission</a></li>
<li><a href="/wiki/Theresa_Wallach" title="Theresa Wallach">Theresa Wallach</a> – Pioneer female engineer, motorcycle adventurer, author, educator and entrepreneur, holder of <a href="/wiki/Brooklands" title="Brooklands">Brooklands</a> Gold Star.</li></ul>
<div class="mw-heading mw-heading4"><h4 id="Business_and_finance">Business and finance</h4><span class="mw-editsection"><span class="mw-editsection-bracket">[</span><a href="/w/index.php?title=City,_University_of_London&amp;action=edit&amp;section=28" title="Edit section: Business and finance"><span>edit</span></a><span class="mw-editsection-bracket">]</span></span></div>
<figure class="mw-default-size mw-halign-right" typeof="mw:File/Thumb"><a href="/wiki/File:Muhtar_A._Kent.jpg" class="mw-file-description"><img src="//upload.wikimedia.org/wikipedia/commons/thumb/7/72/Muhtar_A._Kent.jpg/250px-Muhtar_A._Kent.jpg" decoding="async" width="190" height="127" class="mw-file-element" srcset="//upload.wikimedia.org/wikipedia/commons/thumb/7/72/Muhtar_A._Kent.jpg/330px-Muhtar_A._Kent.jpg 1.5x, //upload.wikimedia.org/wikipedia/commons/thumb/7/72/Muhtar_A._Kent.jpg/500px-Muhtar_A._Kent.jpg 2x" data-file-width="4096" data-file-height="2746" /></a><figcaption><a href="/wiki/Muhtar_Kent" title="Muhtar Kent">Muhtar Kent</a></figcaption></figure>
<ul><li><a href="/wiki/Winston_Set_Aung" class="mw-redirect" title="Winston Set Aung">Winston Set Aung</a> – Politician, Economist and Management Consultant, incumbent Deputy Governor of the <a href="/wiki/Central_Bank_of_Myanmar" title="Central Bank of Myanmar">Central Bank of Myanmar</a></li>
<li><a href="/wiki/Brendan_Barber" title="Brendan Barber">Brendan Barber</a> – former General Secretary of the <a href="/wiki/Trades_Union_Congress" title="Trades Union Congress">Trades Union Congress</a><sup id="cite_ref-52" class="reference"><a href="#cite_note-52"><span class="cite-bracket">&#91;</span>52<span class="cite-bracket">&#93;</span></a></sup></li>
<li>Jonathan Breeze – Founder and CEO of <a href="/wiki/Jet_Republic" title="Jet Republic">Jet Republic</a>, private jet airline company in Europe<sup id="cite_ref-PTC_53-0" class="reference"><a href="#cite_note-PTC-53"><span class="cite-bracket">&#91;</span>53<span class="cite-bracket">&#93;</span></a></sup></li>
<li><a href="/wiki/Michael_Boulos" title="Michael Boulos">Michael Boulos</a> – Associate director of Callian Capital Group, and partner of <a href="/wiki/Tiffany_Trump" title="Tiffany Trump">Tiffany Trump</a><sup id="cite_ref-54" class="reference"><a href="#cite_note-54"><span class="cite-bracket">&#91;</span>54<span class="cite-bracket">&#93;</span></a></sup></li>
<li><a href="/wiki/William_Castell" title="William Castell">William Castell</a> – Former Chairman of the <a href="/wiki/Wellcome_Trust" title="Wellcome Trust">Wellcome Trust</a> and a Director of <a href="/wiki/General_Electric" title="General Electric">General Electric</a> and <a href="/wiki/BP" title="BP">BP</a>, former CEO of <a href="/wiki/Amersham_plc" title="Amersham plc">Amersham plc</a><sup id="cite_ref-PTC_53-1" class="reference"><a href="#cite_note-PTC-53"><span class="cite-bracket">&#91;</span>53<span class="cite-bracket">&#93;</span></a></sup></li>
<li><a href="/wiki/Peter_Cullum" title="Peter Cullum">Peter Cullum</a> – British entrepreneur</li>
<li><a href="/wiki/James_J._Greco" title="James J. Greco">James J. Greco</a> – Former CEO and President of <a href="/wiki/Sbarro" title="Sbarro">Sbarro</a></li>
<li><a href="/wiki/Sir_Stelios_Haji-Ioannou" class="mw-redirect" title="Sir Stelios Haji-Ioannou">Sir Stelios Haji-Ioannou</a> – Founder of <a href="/wiki/EasyGroup" title="EasyGroup">easyGroup</a></li>
<li><a href="/wiki/Tom_Ilube" title="Tom Ilube">Tom Ilube</a> – CBE, British entrepreneur and Chair of the RFU</li>
<li><a href="/wiki/Robert_P._Kelly" title="Robert P. Kelly">Bob Kelly</a> – Former CEO of <a href="/wiki/Bank_of_New_York_Mellon" class="mw-redirect" title="Bank of New York Mellon">Bank of New York Mellon</a> and CFO of <a href="/wiki/Mellon_Financial_Corporation" class="mw-redirect" title="Mellon Financial Corporation">Mellon Financial Corporation</a> and <a href="/wiki/Wachovia_Corporation" class="mw-redirect" title="Wachovia Corporation">Wachovia Corporation</a></li>
<li><a href="/wiki/Muhtar_Kent" title="Muhtar Kent">Muhtar Kent</a> – Former CEO and chairman of <a href="/wiki/The_Coca-Cola_Company" title="The Coca-Cola Company">The Coca-Cola Company</a></li>
<li><a href="/wiki/William_Lewis_(journalist)" title="William Lewis (journalist)">William Lewis</a> – Former CEO <a href="/wiki/Dow_Jones_%26_Company" title="Dow Jones &amp; Company">Dow Jones</a> Publisher, <i>The Wall Street Journal</i><sup id="cite_ref-55" class="reference"><a href="#cite_note-55"><span class="cite-bracket">&#91;</span>55<span class="cite-bracket">&#93;</span></a></sup></li>
<li><a href="/wiki/Ian_Livingstone_(property_developer)" title="Ian Livingstone (property developer)">Ian Livingstone</a> – chairman and co-owner, <a href="/wiki/London_%26_Regional_Properties" title="London &amp; Regional Properties">London &amp; Regional Properties</a><sup id="cite_ref-Questex_56-0" class="reference"><a href="#cite_note-Questex-56"><span class="cite-bracket">&#91;</span>56<span class="cite-bracket">&#93;</span></a></sup></li>
<li><a href="/wiki/Liu_Mingkang" title="Liu Mingkang">Liu Mingkang</a> – Former Chairman of the <a href="/wiki/China_Banking_Regulatory_Commission" title="China Banking Regulatory Commission">China Banking Regulatory Commission</a><sup id="cite_ref-57" class="reference"><a href="#cite_note-57"><span class="cite-bracket">&#91;</span>57<span class="cite-bracket">&#93;</span></a></sup></li>
<li><a href="/wiki/Dick_Olver" title="Dick Olver">Dick Olver</a> – Former Chairman of <a href="/wiki/BAE_Systems" title="BAE Systems">BAE Systems</a>, member of the board of directors at <a href="/wiki/Reuters" title="Reuters">Reuters</a></li>
<li><a href="/wiki/Syed_Ali_Raza" title="Syed Ali Raza">Syed Ali Raza</a> – Former president and Chairman of the <a href="/wiki/National_Bank_of_Pakistan" title="National Bank of Pakistan">National Bank of Pakistan</a></li>
<li><a href="/wiki/Martin_Wheatley" title="Martin Wheatley">Martin Wheatley</a> – Former CEO of the <a href="/wiki/Financial_Conduct_Authority" title="Financial Conduct Authority">Financial Conduct Authority</a></li>
<li><a href="/wiki/Brian_Wynter" title="Brian Wynter">Brian Wynter</a> – Governor of the Bank of Jamaica<sup id="cite_ref-PTC_53-2" class="reference"><a href="#cite_note-PTC-53"><span class="cite-bracket">&#91;</span>53<span class="cite-bracket">&#93;</span></a></sup></li>
<li><a href="/wiki/Durmu%C5%9F_Y%C4%B1lmaz" title="Durmuş Yılmaz">Durmuş Yılmaz</a> – Governor of the Central Bank of Turkey</li></ul>
<div class="mw-heading mw-heading4"><h4 id="Media_and_entertainment">Media and entertainment</h4><span class="mw-editsection"><span class="mw-editsection-bracket">[</span><a href="/w/index.php?title=City,_University_of_London&amp;action=edit&amp;section=29" title="Edit section: Media and entertainment"><span>edit</span></a><span class="mw-editsection-bracket">]</span></span></div>
<figure class="mw-default-size mw-halign-right" typeof="mw:File/Thumb"><a href="/wiki/File:Sophie_raworth.JPG" class="mw-file-description"><img alt="" src="//upload.wikimedia.org/wikipedia/commons/thumb/d/d6/Sophie_raworth.JPG/250px-Sophie_raworth.JPG" decoding="async" width="190" height="285" class="mw-file-element" srcset="//upload.wikimedia.org/wikipedia/commons/thumb/d/d6/Sophie_raworth.JPG/330px-Sophie_raworth.JPG 1.5x, //upload.wikimedia.org/wikipedia/commons/thumb/d/d6/Sophie_raworth.JPG/500px-Sophie_raworth.JPG 2x" data-file-width="2592" data-file-height="3888" /></a><figcaption><a href="/wiki/Sophie_Raworth" title="Sophie Raworth">Sophie Raworth</a></figcaption></figure>
<ul><li><a href="/wiki/Samira_Ahmed" title="Samira Ahmed">Samira Ahmed</a> – <a href="/wiki/Channel_4_News" title="Channel 4 News">Channel 4 News</a> presenter, <a href="/wiki/BBC_News" title="BBC News">BBC News</a> presenter, writer and journalist</li>
<li><a href="/wiki/Decca_Aitkenhead" title="Decca Aitkenhead">Decca Aitkenhead</a> – Journalist<sup id="cite_ref-58" class="reference"><a href="#cite_note-58"><span class="cite-bracket">&#91;</span>58<span class="cite-bracket">&#93;</span></a></sup></li>
<li><a href="/wiki/Joanna_Blythman" title="Joanna Blythman">Joanna Blythman</a> – Non-fiction writer, Britain's leading investigative food journalist<sup id="cite_ref-PTC_53-3" class="reference"><a href="#cite_note-PTC-53"><span class="cite-bracket">&#91;</span>53<span class="cite-bracket">&#93;</span></a></sup></li>
<li><a href="/wiki/Emily_Buchanan" title="Emily Buchanan">Emily Buchanan</a> – BBC World Affairs correspondent</li>
<li><a href="/wiki/Sally_Bundock" title="Sally Bundock">Sally Bundock</a> – BBC presenter</li>
<li><a href="/wiki/Ellie_Crisell" title="Ellie Crisell">Ellie Crisell</a> – <a href="/wiki/BBC" title="BBC">BBC</a> presenter<sup id="cite_ref-PTC_53-4" class="reference"><a href="#cite_note-PTC-53"><span class="cite-bracket">&#91;</span>53<span class="cite-bracket">&#93;</span></a></sup></li>
<li><a href="/wiki/Joe_Crowley_(presenter)" title="Joe Crowley (presenter)">Joe Crowley</a> - Journalist and presenter</li>
<li><a href="/wiki/Imogen_Edwards-Jones" title="Imogen Edwards-Jones">Imogen Edwards-Jones</a><sup class="noprint Inline-Template Template-Fact" style="white-space:nowrap;">&#91;<i><a href="/wiki/Wikipedia:Citation_needed" title="Wikipedia:Citation needed"><span title="This claim needs references to reliable sources. (December 2010)">citation needed</span></a></i>&#93;</sup> – Novelist</li>
<li><a href="/wiki/Gamal_Fahnbulleh" title="Gamal Fahnbulleh">Gamal Fahnbulleh</a> – <a href="/wiki/Sky_News" title="Sky News">Sky News</a> presenter and journalist</li>
<li><a href="/wiki/Mimi_Fawaz" title="Mimi Fawaz">Mimi Fawaz</a>, BBC presenter and journalist<sup id="cite_ref-:02_59-0" class="reference"><a href="#cite_note-:02-59"><span class="cite-bracket">&#91;</span>59<span class="cite-bracket">&#93;</span></a></sup></li>
<li><a href="/wiki/Michael_Fish" title="Michael Fish">Michael Fish</a> – <a href="/wiki/BBC_Weather" title="BBC Weather">BBC</a> weatherman</li>
<li><a href="/wiki/Adam_Fleming_(journalist)" title="Adam Fleming (journalist)">Adam Fleming</a> – <a href="/wiki/CBBC_(TV_channel)" class="mw-redirect" title="CBBC (TV channel)">CBBC</a> reporter</li>
<li><a href="/wiki/Lourdes_Garcia-Navarro" class="mw-redirect" title="Lourdes Garcia-Navarro">Lourdes Garcia-Navarro</a> – Journalist, <a href="/wiki/Jerusalem" title="Jerusalem">Jerusalem</a> foreign correspondent for <a href="/wiki/National_Public_Radio" class="mw-redirect" title="National Public Radio">National Public Radio</a> (NPR)</li>
<li><a href="/wiki/Alex_Graham_(producer)" title="Alex Graham (producer)">Alex Graham</a> – chairman of PACT and the <a href="/wiki/Scott_Trust_Limited" title="Scott Trust Limited">Scott Trust</a></li>
<li><a href="/wiki/Michael_Grothaus" title="Michael Grothaus">Michael Grothaus</a> – Novelist and journalist; author of <i>Epiphany Jones</i></li>
<li><a href="/wiki/Rachel_Horne" title="Rachel Horne">Rachel Horne</a> – BBC and <a href="/wiki/Virgin_Radio" title="Virgin Radio">Virgin Radio</a> presenter and journalist</li>
<li><a href="/wiki/Faisal_Islam" title="Faisal Islam">Faisal Islam</a> – BBC News economics editor</li>
<li><a href="/wiki/Gillian_Joseph" title="Gillian Joseph">Gillian Joseph</a> – Sky News presenter</li>
<li><a href="/wiki/Kirsty_Lang" title="Kirsty Lang">Kirsty Lang</a> – BBC presenter and journalist</li>
<li><a href="/wiki/Ellie_Levenson" title="Ellie Levenson">Ellie Levenson</a> – Freelance journalist and author</li>
<li><a href="/wiki/William_Lewis_(journalist)" title="William Lewis (journalist)">William Lewis</a> – Journalist and editor of <i><a href="/wiki/The_Daily_Telegraph" title="The Daily Telegraph">The Daily Telegraph</a></i></li>
<li><a href="/wiki/Donal_MacIntyre" title="Donal MacIntyre">Donal MacIntyre</a> – Investigative journalist</li>
<li><a href="/wiki/Sharon_Maguire" title="Sharon Maguire">Sharon Maguire</a> – Writer and director, directed <i><a href="/wiki/Bridget_Jones%27s_Diary_(film)" class="mw-redirect" title="Bridget Jones&#39;s Diary (film)">Bridget Jones's Diary</a></i></li>
<li><a href="/wiki/Rhodri_Marsden" title="Rhodri Marsden">Rhodri Marsden</a> – Journalist, musician and blogger; columnist for <i><a href="/wiki/The_Independent" title="The Independent">The Independent</a></i></li>
<li><a href="/wiki/Sharon_Mascall" title="Sharon Mascall">Sharon Mascall</a> – Journalist, broadcaster and writer; lecturer at the <a href="/wiki/University_of_South_Australia" title="University of South Australia">University of South Australia</a></li>
<li><a href="/wiki/Lucrezia_Millarini" title="Lucrezia Millarini">Lucrezia Millarini</a> – Freelance journalist and <a href="/wiki/ITV_(TV_network)" title="ITV (TV network)">ITV</a> newsreader<sup id="cite_ref-60" class="reference"><a href="#cite_note-60"><span class="cite-bracket">&#91;</span>60<span class="cite-bracket">&#93;</span></a></sup></li>
<li><a href="/wiki/Dermot_Murnaghan" title="Dermot Murnaghan">Dermot Murnaghan</a> – Presenter on Sky News</li>
<li><a href="/wiki/Tiff_Needell" title="Tiff Needell">Tiff Needell</a> – <a href="/wiki/Formula_One" title="Formula One">Grand Prix</a> driver, presenter of <i><a href="/wiki/Fifth_Gear" title="Fifth Gear">Fifth Gear</a></i> on <a href="/wiki/Five_(channel)" class="mw-redirect" title="Five (channel)">Five</a></li>
<li><a href="/wiki/Maryam_Nemazee" title="Maryam Nemazee">Maryam Nemazee</a> – presenter for <a href="/wiki/Al_Jazeera_English" title="Al Jazeera English">Al Jazeera</a> London</li>
<li><a href="/wiki/Linda_Papadopoulos" title="Linda Papadopoulos">Linda Papadopoulos</a> – psychologist, appearing occasionally on TV</li>
<li><a href="/wiki/Lilah_Parsons" title="Lilah Parsons">Lilah Parsons</a> - freelance broadcast journalist for <a href="/wiki/BBC_News" title="BBC News">BBC News</a></li>
<li><a href="/wiki/Sebastian_Payne" title="Sebastian Payne">Sebastian Payne</a> – Journalist<sup id="cite_ref-61" class="reference"><a href="#cite_note-61"><span class="cite-bracket">&#91;</span>61<span class="cite-bracket">&#93;</span></a></sup></li>
<li><a href="/wiki/Catherine_Pepinster" title="Catherine Pepinster">Catherine Pepinster</a> – journalist, religion writer</li>
<li><a href="/wiki/Raj_Persaud" title="Raj Persaud">Raj Persaud</a> – British consultant psychiatrist, broadcaster, and author on psychiatry</li>
<li><a href="/wiki/Richard_Preston" title="Richard Preston">Richard Preston</a> – Novelist</li>
<li><a href="/wiki/Gavin_Ramjaun" title="Gavin Ramjaun">Gavin Ramjaun</a> – Television presenter and journalist</li>
<li><a href="/wiki/Sophie_Raworth" title="Sophie Raworth">Sophie Raworth</a> – Newsreader, presenter on BBC One O'Clock News</li>
<li><a href="/wiki/Apsara_Reddy" title="Apsara Reddy">Apsara Reddy</a> – Journalist</li>
<li><a href="/wiki/Joel_Rubin" title="Joel Rubin">Joel Rubin</a> – World-renowned <i>klezmer</i> <a href="/wiki/Clarinet" title="Clarinet">clarinetist</a><sup id="cite_ref-PTC_53-5" class="reference"><a href="#cite_note-PTC-53"><span class="cite-bracket">&#91;</span>53<span class="cite-bracket">&#93;</span></a></sup></li>
<li><a href="/wiki/Ian_Saville" title="Ian Saville">Ian Saville</a> – British magician<sup id="cite_ref-PTC_53-6" class="reference"><a href="#cite_note-PTC-53"><span class="cite-bracket">&#91;</span>53<span class="cite-bracket">&#93;</span></a></sup></li>
<li><a href="/wiki/Barbara_Serra" title="Barbara Serra">Barbara Serra</a> – Presenter for <a href="/wiki/Al_Jazeera_English" title="Al Jazeera English">Al Jazeera</a> London<sup id="cite_ref-PTC_53-7" class="reference"><a href="#cite_note-PTC-53"><span class="cite-bracket">&#91;</span>53<span class="cite-bracket">&#93;</span></a></sup></li>
<li><a href="/wiki/Sarah_Walker_(music_broadcaster)" title="Sarah Walker (music broadcaster)">Sarah Walker</a> – BBC Radio 3 presenter</li>
<li><a href="/wiki/Josh_Widdicombe" title="Josh Widdicombe">Josh Widdicombe</a> – Comedian and presenter<sup id="cite_ref-62" class="reference"><a href="#cite_note-62"><span class="cite-bracket">&#91;</span>62<span class="cite-bracket">&#93;</span></a></sup></li></ul>
<div class="mw-heading mw-heading3"><h3 id="Notable_faculty_and_staff">Notable faculty and staff</h3><span class="mw-editsection"><span class="mw-editsection-bracket">[</span><a href="/w/index.php?title=City,_University_of_London&amp;action=edit&amp;section=30" title="Edit section: Notable faculty and staff"><span>edit</span></a><span class="mw-editsection-bracket">]</span></span></div>
<figure class="mw-default-size mw-halign-right" typeof="mw:File/Thumb"><a href="/wiki/File:David_Willetts.jpg" class="mw-file-description"><img src="//upload.wikimedia.org/wikipedia/commons/thumb/3/33/David_Willetts.jpg/250px-David_Willetts.jpg" decoding="async" width="190" height="258" class="mw-file-element" srcset="//upload.wikimedia.org/wikipedia/commons/thumb/3/33/David_Willetts.jpg/330px-David_Willetts.jpg 1.5x, //upload.wikimedia.org/wikipedia/commons/thumb/3/33/David_Willetts.jpg/500px-David_Willetts.jpg 2x" data-file-width="753" data-file-height="1024" /></a><figcaption><a href="/wiki/David_Willets" class="mw-redirect" title="David Willets">David Willets</a></figcaption></figure>
<ul><li><a href="/wiki/Rosemary_Crompton" title="Rosemary Crompton">Rosemary Crompton</a> –  Professor of Sociology</li>
<li><a href="/wiki/Roy_Greenslade" title="Roy Greenslade">Roy Greenslade</a> – Journalist</li>
<li><a href="/wiki/Steven_Haberman" title="Steven Haberman">Steven Haberman</a> – Professor of <a href="/wiki/Actuarial_Science" class="mw-redirect" title="Actuarial Science">Actuarial Science</a> at City St George's, University of London</li>
<li><a href="/wiki/Corinna_Hawkes" title="Corinna Hawkes">Corinna Hawkes</a> – Professor of <a href="/wiki/Food_policy" title="Food policy">Food Policy</a></li>
<li><a href="/wiki/Rosemary_Hollis" title="Rosemary Hollis">Rosemary Hollis</a> – Professor of <a href="/wiki/International_Politics" class="mw-redirect" title="International Politics">International Politics</a> at City St George's, University of London</li>
<li><a href="/wiki/Jamal_Nazrul_Islam" title="Jamal Nazrul Islam">Jamal Nazrul Islam</a> – Physicist, Mathematician, Cosmologist, Astronomer</li>
<li><a href="/wiki/Ernest_Krausz" title="Ernest Krausz">Ernest Krausz</a> (1931-2018) - Israeli professor of sociology and President at <a href="/wiki/Bar_Ilan_University" class="mw-redirect" title="Bar Ilan University">Bar Ilan University</a></li>
<li><a href="/wiki/David_Leigh_(journalist)" title="David Leigh (journalist)">David Leigh</a> – Journalist</li>
<li><a href="/wiki/David_Marks_(psychologist)" title="David Marks (psychologist)">David Marks</a> – Psychologist</li>
<li><a href="/wiki/Penny_Marshall_(UK_journalist)" class="mw-redirect" title="Penny Marshall (UK journalist)">Penny Marshall</a> – Journalist</li>
<li><a href="/wiki/Stewart_Purvis" title="Stewart Purvis">Stewart Purvis</a> – Broadcaster</li>
<li><a href="/wiki/Denis_Smalley" title="Denis Smalley">Denis Smalley</a> – Composer</li>
<li><a href="/wiki/Bill_Thompson_(technology_writer)" title="Bill Thompson (technology writer)">Bill Thompson</a> – Journalist</li>
<li><a href="/wiki/David_Willets" class="mw-redirect" title="David Willets">David Willets</a> – <a href="/wiki/Conservative_Party_(UK)" title="Conservative Party (UK)">Conservative</a> Member of Parliament for <a href="/wiki/Havant_(constituency)" class="mw-redirect" title="Havant (constituency)">Havant</a>; Shadow <a href="/wiki/Secretary_of_State_for_Education_and_Skills" class="mw-redirect" title="Secretary of State for Education and Skills">Secretary of State for Education and Skills</a></li></ul>
<div class="mw-heading mw-heading3"><h3 id="Vice-Chancellors_(Pre-2016)_/_Presidents_(Post-2016)"><span id="Vice-Chancellors_.28Pre-2016.29_.2F_Presidents_.28Post-2016.29"></span>Vice-Chancellors (Pre-2016) / Presidents (Post-2016)</h3><span class="mw-editsection"><span class="mw-editsection-bracket">[</span><a href="/w/index.php?title=City,_University_of_London&amp;action=edit&amp;section=31" title="Edit section: Vice-Chancellors (Pre-2016) / Presidents (Post-2016)"><span>edit</span></a><span class="mw-editsection-bracket">]</span></span></div>
<ul><li>1966–1974: Sir <a href="/wiki/James_Sharp_Tait" title="James Sharp Tait">James Sharp Tait</a></li>
<li>1974–1978: Sir <a href="/wiki/Edward_Parkes" title="Edward Parkes">Edward W. Parkes</a></li>
<li>1978–1998: <a href="/wiki/Raoul_Franklin" title="Raoul Franklin">Raoul Franklin</a></li>
<li>1998–2007: <a href="/wiki/David_William_Rhind" title="David William Rhind">David William Rhind</a></li>
<li>2007–2009: <a href="/wiki/Malcolm_Gillies" title="Malcolm Gillies">Malcolm Gillies</a></li>
<li>2009–2010: <a href="/wiki/Julius_Weinberg" title="Julius Weinberg">Julius Weinberg</a> (acting)</li>
<li>2010–2021: Sir <a href="/wiki/Paul_Curran_(geographer)" title="Paul Curran (geographer)">Paul Curran</a></li>
<li>2021–Present: Sir <a href="/wiki/Anthony_Finkelstein" title="Anthony Finkelstein">Anthony Finkelstein</a></li></ul>
<div class="mw-heading mw-heading2"><h2 id="In_popular_culture">In popular culture</h2><span class="mw-editsection"><span class="mw-editsection-bracket">[</span><a href="/w/index.php?title=City,_University_of_London&amp;action=edit&amp;section=32" title="Edit section: In popular culture"><span>edit</span></a><span class="mw-editsection-bracket">]</span></span></div>
<p>City University's Bastwick Street Halls of Residence in <a href="/wiki/Islington" title="Islington">Islington</a> was the first home of <i><a href="/wiki/MasterChef_(British_TV_series)" title="MasterChef (British TV series)">MasterChef</a></i> following its 2005 revival.<sup id="cite_ref-63" class="reference"><a href="#cite_note-63"><span class="cite-bracket">&#91;</span>63<span class="cite-bracket">&#93;</span></a></sup><sup id="cite_ref-64" class="reference"><a href="#cite_note-64"><span class="cite-bracket">&#91;</span>64<span class="cite-bracket">&#93;</span></a></sup>
</p>
<div class="mw-heading mw-heading2"><h2 id="See_also">See also</h2><span class="mw-editsection"><span class="mw-editsection-bracket">[</span><a href="/w/index.php?title=City,_University_of_London&amp;action=edit&amp;section=33" title="Edit section: See also"><span>edit</span></a><span class="mw-editsection-bracket">]</span></span></div>
<ul><li><a href="/wiki/Armorial_of_UK_universities" class="mw-redirect" title="Armorial of UK universities">Armorial of UK universities</a></li>
<li><a href="/wiki/College_of_advanced_technology_(United_Kingdom)" title="College of advanced technology (United Kingdom)">College of advanced technology (United Kingdom)</a></li>
<li><a href="/wiki/List_of_universities_in_the_UK" class="mw-redirect" title="List of universities in the UK">List of universities in the UK</a></li></ul>
<div class="mw-heading mw-heading2"><h2 id="References">References</h2><span class="mw-editsection"><span class="mw-editsection-bracket">[</span><a href="/w/index.php?title=City,_University_of_London&amp;action=edit&amp;section=34" title="Edit section: References"><span>edit</span></a><span class="mw-editsection-bracket">]</span></span></div>
<style data-mw-deduplicate="TemplateStyles:r1239543626">.mw-parser-output .reflist{margin-bottom:0.5em;list-style-type:decimal}@media screen{.mw-parser-output .reflist{font-size:90%}}.mw-parser-output .reflist .references{font-size:100%;margin-bottom:0;list-style-type:inherit}.mw-parser-output .reflist-columns-2{column-width:30em}.mw-parser-output .reflist-columns-3{column-width:25em}.mw-parser-output .reflist-columns{margin-top:0.3em}.mw-parser-output .reflist-columns ol{margin-top:0}.mw-parser-output .reflist-columns li{page-break-inside:avoid;break-inside:avoid-column}.mw-parser-output .reflist-upper-alpha{list-style-type:upper-alpha}.mw-parser-output .reflist-upper-roman{list-style-type:upper-roman}.mw-parser-output .reflist-lower-alpha{list-style-type:lower-alpha}.mw-parser-output .reflist-lower-greek{list-style-type:lower-greek}.mw-parser-output .reflist-lower-roman{list-style-type:lower-roman}</style><div class="reflist">
<div class="mw-references-wrap mw-references-columns"><ol class="references">
<li id="cite_note-City_Financial_Statement_23/24-1"><span class="mw-cite-backlink">^ <a href="#cite_ref-City_Financial_Statement_23/24_1-0"><sup><i><b>a</b></i></sup></a> <a href="#cite_ref-City_Financial_Statement_23/24_1-1"><sup><i><b>b</b></i></sup></a></span> <span class="reference-text"><style data-mw-deduplicate="TemplateStyles:r1238218222">.mw-parser-output cite.citation{font-style:inherit;word-wrap:break-word}.mw-parser-output .citation q{quotes:"\"""\"""'""'"}.mw-parser-output .citation:target{background-color:rgba(0,127,255,0.133)}.mw-parser-output .id-lock-free.id-lock-free a{background:url("//upload.wikimedia.org/wikipedia/commons/6/65/Lock-green.svg")right 0.1em center/9px no-repeat}.mw-parser-output .id-lock-limited.id-lock-limited a,.mw-parser-output .id-lock-registration.id-lock-registration a{background:url("//upload.wikimedia.org/wikipedia/commons/d/d6/Lock-gray-alt-2.svg")right 0.1em center/9px no-repeat}.mw-parser-output .id-lock-subscription.id-lock-subscription a{background:url("//upload.wikimedia.org/wikipedia/commons/a/aa/Lock-red-alt-2.svg")right 0.1em center/9px no-repeat}.mw-parser-output .cs1-ws-icon a{background:url("//upload.wikimedia.org/wikipedia/commons/4/4c/Wikisource-logo.svg")right 0.1em center/12px no-repeat}body:not(.skin-timeless):not(.skin-minerva) .mw-parser-output .id-lock-free a,body:not(.skin-timeless):not(.skin-minerva) .mw-parser-output .id-lock-limited a,body:not(.skin-timeless):not(.skin-minerva) .mw-parser-output .id-lock-registration a,body:not(.skin-timeless):not(.skin-minerva) .mw-parser-output .id-lock-subscription a,body:not(.skin-timeless):not(.skin-minerva) .mw-parser-output .cs1-ws-icon a{background-size:contain;padding:0 1em 0 0}.mw-parser-output .cs1-code{color:inherit;background:inherit;border:none;padding:inherit}.mw-parser-output .cs1-hidden-error{display:none;color:var(--color-error,#d33)}.mw-parser-output .cs1-visible-error{color:var(--color-error,#d33)}.mw-parser-output .cs1-maint{display:none;color:#085;margin-left:0.3em}.mw-parser-output .cs1-kern-left{padding-left:0.2em}.mw-parser-output .cs1-kern-right{padding-right:0.2em}.mw-parser-output .citation .mw-selflink{font-weight:inherit}@media screen{.mw-parser-output .cs1-format{font-size:95%}html.skin-theme-clientpref-night .mw-parser-output .cs1-maint{color:#18911f}}@media screen and (prefers-color-scheme:dark){html.skin-theme-clientpref-os .mw-parser-output .cs1-maint{color:#18911f}}</style><cite class="citation web cs1"><a rel="nofollow" class="external text" href="https://www.citystgeorges.ac.uk/__data/assets/pdf_file/0006/845493/Annual-financial-report-2024.pdf">"Financial Statements for the Year to 31 July 2024"</a> <span class="cs1-format">(PDF)</span>. City, University of London<span class="reference-accessdate">. Retrieved <span class="nowrap">14 February</span> 2025</span>.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Abook&amp;rft.genre=unknown&amp;rft.btitle=Financial+Statements+for+the+Year+to+31+July+2024&amp;rft.pub=City%2C+University+of+London&amp;rft_id=https%3A%2F%2Fwww.citystgeorges.ac.uk%2F&#95;_data%2Fassets%2Fpdf_file%2F0006%2F845493%2FAnnual-financial-report-2024.pdf&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
<li id="cite_note-HESAWhere2024-2"><span class="mw-cite-backlink">^ <a href="#cite_ref-HESAWhere2024_2-0"><sup><i><b>a</b></i></sup></a> <a href="#cite_ref-HESAWhere2024_2-1"><sup><i><b>b</b></i></sup></a> <a href="#cite_ref-HESAWhere2024_2-2"><sup><i><b>c</b></i></sup></a></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite class="citation web cs1"><a rel="nofollow" class="external text" href="https://www.hesa.ac.uk/data-and-analysis/students/where-study">"Where do HE students study?"</a>. <i>Higher Education Statistics Agency (HESA) www.hesa.ac.uk</i>. 2024<span class="reference-accessdate">. Retrieved <span class="nowrap">17 September</span> 2024</span>.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Ajournal&amp;rft.genre=unknown&amp;rft.jtitle=Higher+Education+Statistics+Agency+%28HESA%29+www.hesa.ac.uk&amp;rft.atitle=Where+do+HE+students+study%3F&amp;rft.date=2024&amp;rft_id=https%3A%2F%2Fwww.hesa.ac.uk%2Fdata-and-analysis%2Fstudents%2Fwhere-study&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
<li id="cite_note-3"><span class="mw-cite-backlink"><b><a href="#cite_ref-3">^</a></b></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite id="CITEREFChris_Havergal2024" class="citation news cs1">Chris Havergal (22 February 2024). <a rel="nofollow" class="external text" href="https://www.timeshighereducation.com/news/city-and-st-georges-merger-confirmed-summer">"City and St George's merger confirmed for this summer"</a>. <i>Times Higher Education</i>.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Ajournal&amp;rft.genre=article&amp;rft.jtitle=Times+Higher+Education&amp;rft.atitle=City+and+St+George%27s+merger+confirmed+for+this+summer&amp;rft.date=2024-02-22&amp;rft.au=Chris+Havergal&amp;rft_id=https%3A%2F%2Fwww.timeshighereducation.com%2Fnews%2Fcity-and-st-georges-merger-confirmed-summer&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
<li id="cite_note-4"><span class="mw-cite-backlink"><b><a href="#cite_ref-4">^</a></b></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite class="citation web cs1"><a rel="nofollow" class="external text" href="https://www.sgul.ac.uk/news/st-georges-and-city-formally-and-legally-merged-to-form-city-st-georges-university-of-london">"St George's and City formally and legally merged to form City St George's, University of London"</a>. <i>www.sgul.ac.uk</i>. 1 August 2024<span class="reference-accessdate">. Retrieved <span class="nowrap">31 August</span> 2024</span>.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Ajournal&amp;rft.genre=unknown&amp;rft.jtitle=www.sgul.ac.uk&amp;rft.atitle=St+George%27s+and+City+formally+and+legally+merged+to+form+City+St+George%27s%2C+University+of+London&amp;rft.date=2024-08-01&amp;rft_id=https%3A%2F%2Fwww.sgul.ac.uk%2Fnews%2Fst-georges-and-city-formally-and-legally-merged-to-form-city-st-georges-university-of-london&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
<li id="cite_note-5"><span class="mw-cite-backlink"><b><a href="#cite_ref-5">^</a></b></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite class="citation web cs1"><a rel="nofollow" class="external text" href="https://www.citystgeorges.ac.uk/__data/assets/pdf_file/0005/74867/CharterandStatute.pdf">"Royal Charter"</a> <span class="cs1-format">(PDF)</span><span class="reference-accessdate">. Retrieved <span class="nowrap">21 January</span> 2013</span>.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Abook&amp;rft.genre=unknown&amp;rft.btitle=Royal+Charter&amp;rft_id=https%3A%2F%2Fwww.citystgeorges.ac.uk%2F&#95;_data%2Fassets%2Fpdf_file%2F0005%2F74867%2FCharterandStatute.pdf&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
<li id="cite_note-6"><span class="mw-cite-backlink"><b><a href="#cite_ref-6">^</a></b></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite class="citation web cs1"><a rel="nofollow" class="external text" href="https://web.archive.org/web/20130118111911/http://www.city.ac.uk/about/city-information/city-review-2011/a-history-of-city-university-london">"A History of City University London"</a>. City University London. Archived from <a rel="nofollow" class="external text" href="http://www.city.ac.uk/about/city-information/city-review-2011/a-history-of-city-university-london">the original</a> on 18 January 2013<span class="reference-accessdate">. Retrieved <span class="nowrap">1 October</span> 2012</span>.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Abook&amp;rft.genre=unknown&amp;rft.btitle=A+History+of+City+University+London&amp;rft.pub=City+University+London&amp;rft_id=http%3A%2F%2Fwww.city.ac.uk%2Fabout%2Fcity-information%2Fcity-review-2011%2Fa-history-of-city-university-london&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
<li id="cite_note-Grove-7"><span class="mw-cite-backlink">^ <a href="#cite_ref-Grove_7-0"><sup><i><b>a</b></i></sup></a> <a href="#cite_ref-Grove_7-1"><sup><i><b>b</b></i></sup></a></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite id="CITEREFGrove2015" class="citation web cs1">Grove, Jack (16 July 2015). <a rel="nofollow" class="external text" href="https://www.timeshighereducation.co.uk/city-university-london-join-university-london">"City University London to join University of London"</a>. <i>Times Higher Education</i><span class="reference-accessdate">. Retrieved <span class="nowrap">16 July</span> 2015</span>.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Ajournal&amp;rft.genre=unknown&amp;rft.jtitle=Times+Higher+Education&amp;rft.atitle=City+University+London+to+join+University+of+London&amp;rft.date=2015-07-16&amp;rft.aulast=Grove&amp;rft.aufirst=Jack&amp;rft_id=https%3A%2F%2Fwww.timeshighereducation.co.uk%2Fcity-university-london-join-university-london&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
<li id="cite_note-8"><span class="mw-cite-backlink"><b><a href="#cite_ref-8">^</a></b></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite class="citation web cs1"><a rel="nofollow" class="external text" href="https://www.findamasters.com/masters-degrees/browsebydept.aspx?IID=393">"City St George's, University of London"</a>. <i>www.FindAMasters.com</i><span class="reference-accessdate">. Retrieved <span class="nowrap">20 November</span> 2017</span>.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Ajournal&amp;rft.genre=unknown&amp;rft.jtitle=www.FindAMasters.com&amp;rft.atitle=City+St+George%27s%2C+University+of+London&amp;rft_id=https%3A%2F%2Fwww.findamasters.com%2Fmasters-degrees%2Fbrowsebydept.aspx%3FIID%3D393&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
<li id="cite_note-9"><span class="mw-cite-backlink"><b><a href="#cite_ref-9">^</a></b></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite class="citation web cs1"><a rel="nofollow" class="external text" href="https://www.city.ac.uk/about/facts-and-achievements/the-city-of-london-and-the-university">"The City of London and City St George's, University of London"</a>. <i>City St George’s, University of London</i><span class="reference-accessdate">. Retrieved <span class="nowrap">20 November</span> 2017</span>.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Ajournal&amp;rft.genre=unknown&amp;rft.jtitle=City+St+George%E2%80%99s%2C+University+of+London&amp;rft.atitle=The+City+of+London+and+City+St+George%27s%2C+University+of+London&amp;rft_id=https%3A%2F%2Fwww.city.ac.uk%2Fabout%2Ffacts-and-achievements%2Fthe-city-of-london-and-the-university&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
<li id="cite_note-10"><span class="mw-cite-backlink"><b><a href="#cite_ref-10">^</a></b></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite class="citation web cs1"><a rel="nofollow" class="external text" href="https://www.citystgeorges.ac.uk/about/schools">"Academic Schools and departments"</a>. <i>City St George’s, University of London</i><span class="reference-accessdate">. Retrieved <span class="nowrap">7 January</span> 2019</span>.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Ajournal&amp;rft.genre=unknown&amp;rft.jtitle=City+St+George%E2%80%99s%2C+University+of+London&amp;rft.atitle=Academic+Schools+and+departments&amp;rft_id=https%3A%2F%2Fwww.citystgeorges.ac.uk%2Fabout%2Fschools&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
<li id="cite_note-11"><span class="mw-cite-backlink"><b><a href="#cite_ref-11">^</a></b></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite class="citation web cs1"><a rel="nofollow" class="external text" href="https://www.citystgeorges.ac.uk/about/schools">"Schools and Academic Departments"</a>. City University London<span class="reference-accessdate">. Retrieved <span class="nowrap">11 December</span> 2011</span>.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Abook&amp;rft.genre=unknown&amp;rft.btitle=Schools+and+Academic+Departments&amp;rft.pub=City+University+London&amp;rft_id=https%3A%2F%2Fwww.citystgeorges.ac.uk%2Fabout%2Fschools&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
<li id="cite_note-City_Financial_Statement_21/22-12"><span class="mw-cite-backlink"><b><a href="#cite_ref-City_Financial_Statement_21/22_12-0">^</a></b></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite class="citation web cs1"><a rel="nofollow" class="external text" href="https://www.citystgeorges.ac.uk/__data/assets/pdf_file/0020/710435/2639-Annual-financial-report-2022_Full-document_DK1j-accessible.pdf">"Financial Statements for the Year to 31 July 2022"</a> <span class="cs1-format">(PDF)</span>. City, University of London<span class="reference-accessdate">. Retrieved <span class="nowrap">14 February</span> 2025</span>.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Abook&amp;rft.genre=unknown&amp;rft.btitle=Financial+Statements+for+the+Year+to+31+July+2022&amp;rft.pub=City%2C+University+of+London&amp;rft_id=https%3A%2F%2Fwww.citystgeorges.ac.uk%2F&#95;_data%2Fassets%2Fpdf_file%2F0020%2F710435%2F2639-Annual-financial-report-2022_Full-document_DK1j-accessible.pdf&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
<li id="cite_note-:1-13"><span class="mw-cite-backlink">^ <a href="#cite_ref-:1_13-0"><sup><i><b>a</b></i></sup></a> <a href="#cite_ref-:1_13-1"><sup><i><b>b</b></i></sup></a></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite class="citation web cs1"><a rel="nofollow" class="external text" href="https://web.archive.org/web/20150717092334/http://www.city.ac.uk/international/international-partnerships/wc2-university-network">"WC2 University Network"</a>. <i>City University London</i>. Archived from <a rel="nofollow" class="external text" href="http://www.city.ac.uk/international/international-partnerships/wc2-university-network">the original</a> on 17 July 2015<span class="reference-accessdate">. Retrieved <span class="nowrap">16 July</span> 2015</span>.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Ajournal&amp;rft.genre=unknown&amp;rft.jtitle=City+University+London&amp;rft.atitle=WC2+University+Network&amp;rft_id=http%3A%2F%2Fwww.city.ac.uk%2Finternational%2Finternational-partnerships%2Fwc2-university-network&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
<li id="cite_note-14"><span class="mw-cite-backlink"><b><a href="#cite_ref-14">^</a></b></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite class="citation web cs1"><a rel="nofollow" class="external text" href="https://www.citystgeorges.ac.uk/about/schools/law/equality-diversity-and-inclusion">"Equality, Diversity and Inclusion at The City Law School | City St George's, University of London"</a>. <i>www.citystgeorges.ac.uk</i>. 6 December 2022<span class="reference-accessdate">. Retrieved <span class="nowrap">11 April</span> 2024</span>.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Ajournal&amp;rft.genre=unknown&amp;rft.jtitle=www.citystgeorges.ac.uk&amp;rft.atitle=Equality%2C+Diversity+and+Inclusion+at+The+City+Law+School+%7C+City+St+George%27s%2C+University+of+London&amp;rft.date=2022-12-06&amp;rft_id=https%3A%2F%2Fwww.citystgeorges.ac.uk%2Fabout%2Fschools%2Flaw%2Fequality-diversity-and-inclusion&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
<li id="cite_note-CityHistory-15"><span class="mw-cite-backlink">^ <a href="#cite_ref-CityHistory_15-0"><sup><i><b>a</b></i></sup></a> <a href="#cite_ref-CityHistory_15-1"><sup><i><b>b</b></i></sup></a> <a href="#cite_ref-CityHistory_15-2"><sup><i><b>c</b></i></sup></a> <a href="#cite_ref-CityHistory_15-3"><sup><i><b>d</b></i></sup></a></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite class="citation web cs1"><a rel="nofollow" class="external text" href="https://web.archive.org/web/20080111154506/https://www.citystgeorges.ac.uk/about/history">"Our history – City University London"</a>. City University, London. Archived from <a rel="nofollow" class="external text" href="https://www.citystgeorges.ac.uk/about/history">the original</a> on 11 January 2008<span class="reference-accessdate">. Retrieved <span class="nowrap">30 January</span> 2008</span>.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Abook&amp;rft.genre=unknown&amp;rft.btitle=Our+history+%E2%80%93+City+University+London&amp;rft.pub=City+University%2C+London&amp;rft_id=https%3A%2F%2Fwww.citystgeorges.ac.uk%2Fabout%2Fhistory&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
<li id="cite_note-16"><span class="mw-cite-backlink"><b><a href="#cite_ref-16">^</a></b></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite class="citation web cs1"><a rel="nofollow" class="external text" href="https://125-anniversary.city.ac.uk/the-pioneering-principal/">"The pioneering Principal"</a>. <i>City 125th Anniversary</i>. 18 November 2018.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Ajournal&amp;rft.genre=unknown&amp;rft.jtitle=City+125th+Anniversary&amp;rft.atitle=The+pioneering+Principal&amp;rft.date=2018-11-18&amp;rft_id=https%3A%2F%2F125-anniversary.city.ac.uk%2Fthe-pioneering-principal%2F&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
<li id="cite_note-17"><span class="mw-cite-backlink"><b><a href="#cite_ref-17">^</a></b></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite class="citation web cs1"><a rel="nofollow" class="external text" href="https://web.archive.org/web/20100914105055/http://www.shl.lon.ac.uk/specialcollections/archives/studentrecords.shtml">"University of London Students 1836–1933"</a>. Senate House Library. 30 June 1930. Archived from <a rel="nofollow" class="external text" href="http://www.senatehouselibrary.ac.uk/our-collections/historic-collections/archives-manuscripts/university-of-london-student-records-1836-1931/">the original</a> on 14 September 2010<span class="reference-accessdate">. Retrieved <span class="nowrap">15 November</span> 2013</span>.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Abook&amp;rft.genre=unknown&amp;rft.btitle=University+of+London+Students+1836%E2%80%931933&amp;rft.pub=Senate+House+Library&amp;rft.date=1930-06-30&amp;rft_id=http%3A%2F%2Fwww.senatehouselibrary.ac.uk%2Four-collections%2Fhistoric-collections%2Farchives-manuscripts%2Funiversity-of-london-student-records-1836-1931%2F&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
<li id="cite_note-18"><span class="mw-cite-backlink"><b><a href="#cite_ref-18">^</a></b></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite class="citation web cs1"><a rel="nofollow" class="external text" href="http://www.city.ac.uk/sems/dps/events/programme.pdf">"100 years of education in aeronautics"</a> <span class="cs1-format">(PDF)</span>. Royal Aeronautical Society<span class="reference-accessdate">. Retrieved <span class="nowrap">15 June</span> 2009</span>.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Abook&amp;rft.genre=unknown&amp;rft.btitle=100+years+of+education+in+aeronautics&amp;rft.pub=Royal+Aeronautical+Society&amp;rft_id=http%3A%2F%2Fwww.city.ac.uk%2Fsems%2Fdps%2Fevents%2Fprogramme.pdf&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span><sup class="noprint Inline-Template"><span style="white-space: nowrap;">&#91;<i><a href="/wiki/Wikipedia:Link_rot" title="Wikipedia:Link rot"><span title="&#160;Dead link tagged November 2017">permanent dead link</span></a></i>&#93;</span></sup></span>
</li>
<li id="cite_note-19"><span class="mw-cite-backlink"><b><a href="#cite_ref-19">^</a></b></span> <span class="reference-text"><a rel="nofollow" class="external text" href="https://web.archive.org/web/20070927222443/http://www.la84foundation.org/6oic/OfficialReports/1908/1908.pdf">1908 Summer Olympics official report.</a> p 33.</span>
</li>
<li id="cite_note-20"><span class="mw-cite-backlink"><b><a href="#cite_ref-20">^</a></b></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite class="citation web cs1"><a rel="nofollow" class="external text" href="https://www.webcitation.org/5rYWTZ2dk?url=http://www.city.ac.uk/aboutcity/dps/Recent-History-Of-City-University-London-APO-Williams.pdf">"Progressing through change: The Recent History of City University London, 1978–2008"</a> <span class="cs1-format">(PDF)</span>. Archived from <a rel="nofollow" class="external text" href="http://www.city.ac.uk/aboutcity/dps/Recent-History-Of-City-University-London-APO-Williams.pdf">the original</a> <span class="cs1-format">(PDF)</span> on 28 July 2010<span class="reference-accessdate">. Retrieved <span class="nowrap">17 July</span> 2009</span>.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Abook&amp;rft.genre=unknown&amp;rft.btitle=Progressing+through+change%3A+The+Recent+History+of+City+University+London%2C+1978%E2%80%932008&amp;rft_id=http%3A%2F%2Fwww.city.ac.uk%2Faboutcity%2Fdps%2FRecent-History-Of-City-University-London-APO-Williams.pdf&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
<li id="cite_note-21"><span class="mw-cite-backlink"><b><a href="#cite_ref-21">^</a></b></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite class="citation web cs1"><a rel="nofollow" class="external text" href="https://web.archive.org/web/20090911091520/http://www.city.ac.uk/sems/apollo-15-astronauts.html">"Video of Apollo 15 astronauts visiting City"</a>. City University London. Archived from <a rel="nofollow" class="external text" href="http://www.city.ac.uk/sems/apollo-15-astronauts.html">the original</a> on 11 September 2009<span class="reference-accessdate">. Retrieved <span class="nowrap">15 June</span> 2009</span>.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Abook&amp;rft.genre=unknown&amp;rft.btitle=Video+of+Apollo+15+astronauts+visiting+City&amp;rft.pub=City+University+London&amp;rft_id=http%3A%2F%2Fwww.city.ac.uk%2Fsems%2Fapollo-15-astronauts.html&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
<li id="cite_note-22"><span class="mw-cite-backlink"><b><a href="#cite_ref-22">^</a></b></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite class="citation news cs1"><a rel="nofollow" class="external text" href="http://www.timeshighereducation.co.uk/story.asp?storyCode=95369&amp;sectioncode=26">"Institute nurses health"</a>. Times Higher Education. 6 October 1995<span class="reference-accessdate">. Retrieved <span class="nowrap">31 December</span> 2011</span>.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Ajournal&amp;rft.genre=article&amp;rft.atitle=Institute+nurses+health&amp;rft.date=1995-10-06&amp;rft_id=http%3A%2F%2Fwww.timeshighereducation.co.uk%2Fstory.asp%3FstoryCode%3D95369%26sectioncode%3D26&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
<li id="cite_note-23"><span class="mw-cite-backlink"><b><a href="#cite_ref-23">^</a></b></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite class="citation news cs1"><a rel="nofollow" class="external text" href="http://www.timeshighereducation.co.uk/story.asp?storyCode=158856&amp;sectioncode=26">"Queen Mary, City kick off alliance"</a>. Times Higher Education. 12 April 2001<span class="reference-accessdate">. Retrieved <span class="nowrap">31 December</span> 2011</span>.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Ajournal&amp;rft.genre=article&amp;rft.atitle=Queen+Mary%2C+City+kick+off+alliance&amp;rft.date=2001-04-12&amp;rft_id=http%3A%2F%2Fwww.timeshighereducation.co.uk%2Fstory.asp%3FstoryCode%3D158856%26sectioncode%3D26&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
<li id="cite_note-24"><span class="mw-cite-backlink"><b><a href="#cite_ref-24">^</a></b></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite id="CITEREFPlomin2001" class="citation news cs1">Plomin, Joe (22 May 2001). <a rel="nofollow" class="external text" href="https://www.theguardian.com/education/2001/may/22/highereducation.news">"Fire destroys part of City University building"</a>. <i>The Guardian</i>. London<span class="reference-accessdate">. Retrieved <span class="nowrap">29 December</span> 2011</span>.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Ajournal&amp;rft.genre=article&amp;rft.jtitle=The+Guardian&amp;rft.atitle=Fire+destroys+part+of+City+University+building&amp;rft.date=2001-05-22&amp;rft.aulast=Plomin&amp;rft.aufirst=Joe&amp;rft_id=https%3A%2F%2Fwww.theguardian.com%2Feducation%2F2001%2Fmay%2F22%2Fhighereducation.news&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
<li id="cite_note-25"><span class="mw-cite-backlink"><b><a href="#cite_ref-25">^</a></b></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite class="citation news cs1"><a rel="nofollow" class="external text" href="http://www.timeshighereducation.co.uk/story.asp?storyCode=164514&amp;sectioncode=26">"Law school to merge with City"</a>. Times Higher Education. 24 August 2001<span class="reference-accessdate">. Retrieved <span class="nowrap">29 December</span> 2011</span>.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Ajournal&amp;rft.genre=article&amp;rft.atitle=Law+school+to+merge+with+City&amp;rft.date=2001-08-24&amp;rft_id=http%3A%2F%2Fwww.timeshighereducation.co.uk%2Fstory.asp%3FstoryCode%3D164514%26sectioncode%3D26&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
<li id="cite_note-26"><span class="mw-cite-backlink"><b><a href="#cite_ref-26">^</a></b></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite class="citation news cs1"><a rel="nofollow" class="external text" href="http://www.timeshighereducation.co.uk/story.asp?storyCode=159680&amp;sectioncode=26">"City Business School seeks global profile"</a>. Times Higher Education. 11 May 2001<span class="reference-accessdate">. Retrieved <span class="nowrap">29 December</span> 2011</span>.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Ajournal&amp;rft.genre=article&amp;rft.atitle=City+Business+School+seeks+global+profile&amp;rft.date=2001-05-11&amp;rft_id=http%3A%2F%2Fwww.timeshighereducation.co.uk%2Fstory.asp%3FstoryCode%3D159680%26sectioncode%3D26&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
<li id="cite_note-27"><span class="mw-cite-backlink"><b><a href="#cite_ref-27">^</a></b></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite class="citation web cs1"><a rel="nofollow" class="external text" href="https://www.citystgeorges.ac.uk/about/history">"Key milestones in our history"</a>. <i>citystgeorges.ac.uk</i>. 23 November 2020.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Ajournal&amp;rft.genre=unknown&amp;rft.jtitle=citystgeorges.ac.uk&amp;rft.atitle=Key+milestones+in+our+history&amp;rft.date=2020-11-23&amp;rft_id=https%3A%2F%2Fwww.citystgeorges.ac.uk%2Fabout%2Fhistory&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
<li id="cite_note-28"><span class="mw-cite-backlink"><b><a href="#cite_ref-28">^</a></b></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite class="citation web cs1"><a rel="nofollow" class="external text" href="https://community.city.ac.uk/city/emailviewonwebpage.aspx?erid=15633926&amp;trid=0db8b50e-31df-43fe-a67c-87b8d5c2b475">"Have your say on the University's new name"</a>. City, University of London. 31 May 2023.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Abook&amp;rft.genre=unknown&amp;rft.btitle=Have+your+say+on+the+University%27s+new+name&amp;rft.pub=City%2C+University+of+London&amp;rft.date=2023-05-31&amp;rft_id=https%3A%2F%2Fcommunity.city.ac.uk%2Fcity%2Femailviewonwebpage.aspx%3Ferid%3D15633926%26trid%3D0db8b50e-31df-43fe-a67c-87b8d5c2b475&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
<li id="cite_note-29"><span class="mw-cite-backlink"><b><a href="#cite_ref-29">^</a></b></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite class="citation web cs1"><a rel="nofollow" class="external text" href="https://www.citystgeorges.ac.uk/about/merger-st-georges">"City, University of London and St George's, University of London officially merge"</a>. <i>www.citystgeorges.ac.uk</i>. 15 August 2024<span class="reference-accessdate">. Retrieved <span class="nowrap">15 August</span> 2024</span>.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Ajournal&amp;rft.genre=unknown&amp;rft.jtitle=www.citystgeorges.ac.uk&amp;rft.atitle=City%2C+University+of+London+and+St+George%27s%2C+University+of+London+officially+merge&amp;rft.date=2024-08-15&amp;rft_id=https%3A%2F%2Fwww.citystgeorges.ac.uk%2Fabout%2Fmerger-st-georges&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
<li id="cite_note-30"><span class="mw-cite-backlink"><b><a href="#cite_ref-30">^</a></b></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite class="citation web cs1"><a rel="nofollow" class="external text" href="https://web.archive.org/web/20090429012107/http://www.city.ac.uk/maps/areamaps/index.html">"University location maps"</a>. City University London. Archived from <a rel="nofollow" class="external text" href="http://www.city.ac.uk/maps/areamaps/index.html">the original</a> on 29 April 2009<span class="reference-accessdate">. Retrieved <span class="nowrap">15 June</span> 2009</span>.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Abook&amp;rft.genre=unknown&amp;rft.btitle=University+location+maps&amp;rft.pub=City+University+London&amp;rft_id=http%3A%2F%2Fwww.city.ac.uk%2Fmaps%2Fareamaps%2Findex.html&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
<li id="cite_note-annrep-31"><span class="mw-cite-backlink">^ <a href="#cite_ref-annrep_31-0"><sup><i><b>a</b></i></sup></a> <a href="#cite_ref-annrep_31-1"><sup><i><b>b</b></i></sup></a> <a href="#cite_ref-annrep_31-2"><sup><i><b>c</b></i></sup></a> <a href="#cite_ref-annrep_31-3"><sup><i><b>d</b></i></sup></a></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite class="citation web cs1"><a rel="nofollow" class="external text" href="http://www.city.ac.uk/__data/assets/pdf_file/0007/110779/Finance_Statements_2011_Web.pdf">"Financial Statements for the year ended 31 July 2011"</a> <span class="cs1-format">(PDF)</span>. City University London<span class="reference-accessdate">. Retrieved <span class="nowrap">24 February</span> 2012</span>.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Abook&amp;rft.genre=unknown&amp;rft.btitle=Financial+Statements+for+the+year+ended+31+July+2011&amp;rft.pub=City+University+London&amp;rft_id=http%3A%2F%2Fwww.city.ac.uk%2F&#95;_data%2Fassets%2Fpdf_file%2F0007%2F110779%2FFinance_Statements_2011_Web.pdf&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
<li id="cite_note-Complete_League_Table-32"><span class="mw-cite-backlink"><b><a href="#cite_ref-Complete_League_Table_32-0">^</a></b></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite class="citation web cs1"><a rel="nofollow" class="external text" href="https://www.thecompleteuniversityguide.co.uk/league-tables/rankings">"Complete University Guide 2025"</a>. The Complete University Guide. 14 May 2024.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Abook&amp;rft.genre=unknown&amp;rft.btitle=Complete+University+Guide+2025&amp;rft.pub=The+Complete+University+Guide&amp;rft.date=2024-05-14&amp;rft_id=https%3A%2F%2Fwww.thecompleteuniversityguide.co.uk%2Fleague-tables%2Frankings&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
<li id="cite_note-The_Guardian_University_Guide-33"><span class="mw-cite-backlink"><b><a href="#cite_ref-The_Guardian_University_Guide_33-0">^</a></b></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite class="citation web cs1"><a rel="nofollow" class="external text" href="https://www.theguardian.com/education/ng-interactive/2024/sep/07/the-guardian-university-guide-2025-the-rankings">"Guardian University Guide 2025"</a>. <i>The Guardian</i>. 7 September 2024.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Ajournal&amp;rft.genre=unknown&amp;rft.jtitle=The+Guardian&amp;rft.atitle=Guardian+University+Guide+2025&amp;rft.date=2024-09-07&amp;rft_id=https%3A%2F%2Fwww.theguardian.com%2Feducation%2Fng-interactive%2F2024%2Fsep%2F07%2Fthe-guardian-university-guide-2025-the-rankings&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
<li id="cite_note-The_Times_and_Sunday_Times_University_Guide-34"><span class="mw-cite-backlink"><b><a href="#cite_ref-The_Times_and_Sunday_Times_University_Guide_34-0">^</a></b></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite class="citation web cs1"><a rel="nofollow" class="external text" href="https://www.thetimes.com/uk-university-rankings/league-table">"Good University Guide 2025"</a>. <i>The Times</i>. 20 September 2024.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Ajournal&amp;rft.genre=unknown&amp;rft.jtitle=The+Times&amp;rft.atitle=Good+University+Guide+2025&amp;rft.date=2024-09-20&amp;rft_id=https%3A%2F%2Fwww.thetimes.com%2Fuk-university-rankings%2Fleague-table&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
<li id="cite_note-Academic_Ranking_of_World_Universities-35"><span class="mw-cite-backlink"><b><a href="#cite_ref-Academic_Ranking_of_World_Universities_35-0">^</a></b></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite class="citation web cs1"><a rel="nofollow" class="external text" href="https://www.shanghairanking.com/rankings/arwu/2024">"Academic Ranking of World Universities 2024"</a>. Shanghai Ranking Consultancy. 15 August 2024.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Abook&amp;rft.genre=unknown&amp;rft.btitle=Academic+Ranking+of+World+Universities+2024&amp;rft.pub=Shanghai+Ranking+Consultancy&amp;rft.date=2024-08-15&amp;rft_id=https%3A%2F%2Fwww.shanghairanking.com%2Frankings%2Farwu%2F2024&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
<li id="cite_note-QS_World_University_Rankings-36"><span class="mw-cite-backlink"><b><a href="#cite_ref-QS_World_University_Rankings_36-0">^</a></b></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite class="citation web cs1"><a rel="nofollow" class="external text" href="https://www.topuniversities.com/world-university-rankings">"QS World University Rankings 2025"</a>. Quacquarelli Symonds Ltd. 4 June 2024.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Abook&amp;rft.genre=unknown&amp;rft.btitle=QS+World+University+Rankings+2025&amp;rft.pub=Quacquarelli+Symonds+Ltd.&amp;rft.date=2024-06-04&amp;rft_id=https%3A%2F%2Fwww.topuniversities.com%2Fworld-university-rankings&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
<li id="cite_note-THE_World_University_Rankings-37"><span class="mw-cite-backlink"><b><a href="#cite_ref-THE_World_University_Rankings_37-0">^</a></b></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite class="citation web cs1"><a rel="nofollow" class="external text" href="https://www.timeshighereducation.com/world-university-rankings/2025/world-ranking">"THE World University Rankings 2025"</a>. Times Higher Education. 9 October 2024.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Abook&amp;rft.genre=unknown&amp;rft.btitle=THE+World+University+Rankings+2025&amp;rft.pub=Times+Higher+Education&amp;rft.date=2024-10-09&amp;rft_id=https%3A%2F%2Fwww.timeshighereducation.com%2Fworld-university-rankings%2F2025%2Fworld-ranking&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
<li id="cite_note-38"><span class="mw-cite-backlink"><b><a href="#cite_ref-38">^</a></b></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite class="citation web cs1"><a rel="nofollow" class="external text" href="https://web.archive.org/web/20090422231507/http://www.city.ac.uk/alumni/careers/index.html">"City's Online Careers Network"</a>. City University London. Archived from <a rel="nofollow" class="external text" href="http://www.city.ac.uk/alumni/careers/index.html">the original</a> on 22 April 2009<span class="reference-accessdate">. Retrieved <span class="nowrap">11 June</span> 2009</span>.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Abook&amp;rft.genre=unknown&amp;rft.btitle=City%27s+Online+Careers+Network&amp;rft.pub=City+University+London&amp;rft_id=http%3A%2F%2Fwww.city.ac.uk%2Falumni%2Fcareers%2Findex.html&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
<li id="cite_note-39"><span class="mw-cite-backlink"><b><a href="#cite_ref-39">^</a></b></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite class="citation web cs1"><a rel="nofollow" class="external text" href="https://web.archive.org/web/20101207151511/http://www.sra.org.uk/documents/students/lpc/exec3innsofcourt.pdf">"Solicitors Regulation Authority Executive Summary"</a> <span class="cs1-format">(PDF)</span>. Solicitors Regulation Authority. 20 March 2007. Archived from <a rel="nofollow" class="external text" href="http://www.sra.org.uk/documents/students/lpc/exec3innsofcourt.pdf">the original</a> <span class="cs1-format">(PDF)</span> on 7 December 2010<span class="reference-accessdate">. Retrieved <span class="nowrap">11 June</span> 2009</span>.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Abook&amp;rft.genre=unknown&amp;rft.btitle=Solicitors+Regulation+Authority+Executive+Summary&amp;rft.pub=Solicitors+Regulation+Authority&amp;rft.date=2007-03-20&amp;rft_id=http%3A%2F%2Fwww.sra.org.uk%2Fdocuments%2Fstudents%2Flpc%2Fexec3innsofcourt.pdf&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
<li id="cite_note-40"><span class="mw-cite-backlink"><b><a href="#cite_ref-40">^</a></b></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite class="citation web cs1"><a rel="nofollow" class="external text" href="http://www.cetl.org.uk/cetl_background.php">"CETL – Centre for Excellence in Teaching and Learning"</a>. Queen Mary University of London<span class="reference-accessdate">. Retrieved <span class="nowrap">25 February</span> 2009</span>.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Abook&amp;rft.genre=unknown&amp;rft.btitle=CETL+%E2%80%93+Centre+for+Excellence+in+Teaching+and+Learning&amp;rft.pub=Queen+Mary+University+of+London&amp;rft_id=http%3A%2F%2Fwww.cetl.org.uk%2Fcetl_background.php&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
<li id="cite_note-:0-41"><span class="mw-cite-backlink"><b><a href="#cite_ref-:0_41-0">^</a></b></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite class="citation web cs1"><a rel="nofollow" class="external text" href="https://web.archive.org/web/20090428031933/http://www.qaa.ac.uk/reviews/reports/subjectlevel/q96_94_textonly.htm">"Links with businesses"</a>. QAA. 24 May 2005. Archived from <a rel="nofollow" class="external text" href="http://www.qaa.ac.uk/reviews/reports/subjectlevel/q96_94_textonly.htm">the original</a> on 28 April 2009<span class="reference-accessdate">. Retrieved <span class="nowrap">21 December</span> 2007</span>.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Abook&amp;rft.genre=unknown&amp;rft.btitle=Links+with+businesses&amp;rft.pub=QAA&amp;rft.date=2005-05-24&amp;rft_id=http%3A%2F%2Fwww.qaa.ac.uk%2Freviews%2Freports%2Fsubjectlevel%2Fq96_94_textonly.htm&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
<li id="cite_note-42"><span class="mw-cite-backlink"><b><a href="#cite_ref-42">^</a></b></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite class="citation web cs1"><a rel="nofollow" class="external text" href="https://web.archive.org/web/20111104055637/http://www.city.ac.uk/news/2011/august/city-university-london-joins-leading-health-research-partnership">"City University London joins leading health research partnership &#124; City University London"</a>. Archived from <a rel="nofollow" class="external text" href="http://www.city.ac.uk/news/2011/august/city-university-london-joins-leading-health-research-partnership">the original</a> on 4 November 2011<span class="reference-accessdate">. Retrieved <span class="nowrap">2 October</span> 2011</span>.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Abook&amp;rft.genre=unknown&amp;rft.btitle=City+University+London+joins+leading+health+research+partnership+%26%23124%3B+City+University+London&amp;rft_id=http%3A%2F%2Fwww.city.ac.uk%2Fnews%2F2011%2Faugust%2Fcity-university-london-joins-leading-health-research-partnership&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
<li id="cite_note-openaccess.city.ac.uk/information-43"><span class="mw-cite-backlink">^ <a href="#cite_ref-openaccess.city.ac.uk/information_43-0"><sup><i><b>a</b></i></sup></a> <a href="#cite_ref-openaccess.city.ac.uk/information_43-1"><sup><i><b>b</b></i></sup></a></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite id="CITEREFCity_Research_Online" class="citation web cs1">City Research Online. <a rel="nofollow" class="external text" href="https://openaccess.city.ac.uk/information.html">"About"</a>. <i>openaccess.city.ac.uk</i><span class="reference-accessdate">. Retrieved <span class="nowrap">11 August</span> 2023</span>.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Ajournal&amp;rft.genre=unknown&amp;rft.jtitle=openaccess.city.ac.uk&amp;rft.atitle=About&amp;rft.au=City+Research+Online&amp;rft_id=https%3A%2F%2Fopenaccess.city.ac.uk%2Finformation.html&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
<li id="cite_note-44"><span class="mw-cite-backlink"><b><a href="#cite_ref-44">^</a></b></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite class="citation web cs1"><a rel="nofollow" class="external text" href="https://web.archive.org/web/20110321083722/http://www.city.ac.uk/studentcentre/studentsunion/about_the_su.html">"About the Students' Union - City University London"</a>. Archived from <a rel="nofollow" class="external text" href="https://www.citystudents.co.uk/about-us/">the original</a> on 21 March 2011.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Abook&amp;rft.genre=unknown&amp;rft.btitle=About+the+Students%27+Union+-+City+University+London&amp;rft_id=https%3A%2F%2Fwww.citystudents.co.uk%2Fabout-us%2F&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
<li id="cite_note-45"><span class="mw-cite-backlink"><b><a href="#cite_ref-45">^</a></b></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite class="citation web cs1"><a rel="nofollow" class="external text" href="https://www.citystudents.co.uk/news/article/6013/Carrot-Radio-to-go-On-Air-soon/">"Carrot Radio to go "On Air" soon"</a>. <i>City Students' Union</i>.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Ajournal&amp;rft.genre=unknown&amp;rft.jtitle=City+Students%27+Union&amp;rft.atitle=Carrot+Radio+to+go+%22On+Air%22+soon&amp;rft_id=https%3A%2F%2Fwww.citystudents.co.uk%2Fnews%2Farticle%2F6013%2FCarrot-Radio-to-go-On-Air-soon%2F&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
<li id="cite_note-46"><span class="mw-cite-backlink"><b><a href="#cite_ref-46">^</a></b></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite id="CITEREFLightfoot2019" class="citation news cs1">Lightfoot, Liz (16 July 2019). <a rel="nofollow" class="external text" href="https://www.theguardian.com/education/2019/jul/16/university-green-rankins-risk-despite-climate-emergency">"University green rankings at risk despite climate emergency"</a>. <i><a href="/wiki/The_Guardian" title="The Guardian">The Guardian</a></i><span class="reference-accessdate">. Retrieved <span class="nowrap">6 July</span> 2020</span>.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Ajournal&amp;rft.genre=article&amp;rft.jtitle=The+Guardian&amp;rft.atitle=University+green+rankings+at+risk+despite+climate+emergency&amp;rft.date=2019-07-16&amp;rft.aulast=Lightfoot&amp;rft.aufirst=Liz&amp;rft_id=https%3A%2F%2Fwww.theguardian.com%2Feducation%2F2019%2Fjul%2F16%2Funiversity-green-rankins-risk-despite-climate-emergency&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
<li id="cite_note-47"><span class="mw-cite-backlink"><b><a href="#cite_ref-47">^</a></b></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite class="citation web cs1"><a rel="nofollow" class="external text" href="https://web.archive.org/web/20170730233624/https://peopleandplanet.org/university/129364/ul16">"City, University of London People &amp; Planet University League 2016 Scorecard"</a>. <a href="/wiki/People_%26_Planet" title="People &amp; Planet">People &amp; Planet</a>. Archived from <a rel="nofollow" class="external text" href="https://peopleandplanet.org/university/129364/ul16">the original</a> on 30 July 2017<span class="reference-accessdate">. Retrieved <span class="nowrap">30 July</span> 2017</span>.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Abook&amp;rft.genre=unknown&amp;rft.btitle=City%2C+University+of+London+People+%26+Planet+University+League+2016+Scorecard&amp;rft.pub=People+%26+Planet&amp;rft_id=https%3A%2F%2Fpeopleandplanet.org%2Funiversity%2F129364%2Ful16&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
<li id="cite_note-48"><span class="mw-cite-backlink"><b><a href="#cite_ref-48">^</a></b></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite class="citation web cs1"><a rel="nofollow" class="external text" href="https://www.citystgeorges.ac.uk/news-and-events/news/2023/07/city-university-of-london-divests-fossil-fuel-producers">"City, University of London divests from fossil fuel producers"</a>. City, University of London. 4 July 2023<span class="reference-accessdate">. Retrieved <span class="nowrap">10 July</span> 2023</span>.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Abook&amp;rft.genre=unknown&amp;rft.btitle=City%2C+University+of+London+divests+from+fossil+fuel+producers&amp;rft.pub=City%2C+University+of+London&amp;rft.date=2023-07-04&amp;rft_id=https%3A%2F%2Fwww.citystgeorges.ac.uk%2Fnews-and-events%2Fnews%2F2023%2F07%2Fcity-university-of-london-divests-fossil-fuel-producers&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
<li id="cite_note-bbc-49"><span class="mw-cite-backlink"><b><a href="#cite_ref-bbc_49-0">^</a></b></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite class="citation news cs1"><a rel="nofollow" class="external text" href="http://news.bbc.co.uk/2/hi/uk_news/7935741.stm">"Broadening Britain's judicial ranks"</a>. <i>BBC</i>. 11 March 2009<span class="reference-accessdate">. Retrieved <span class="nowrap">8 January</span> 2021</span>.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Ajournal&amp;rft.genre=article&amp;rft.jtitle=BBC&amp;rft.atitle=Broadening+Britain%27s+judicial+ranks&amp;rft.date=2009-03-11&amp;rft_id=http%3A%2F%2Fnews.bbc.co.uk%2F2%2Fhi%2Fuk_news%2F7935741.stm&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
<li id="cite_note-50"><span class="mw-cite-backlink"><b><a href="#cite_ref-50">^</a></b></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite class="citation web cs1"><a rel="nofollow" class="external text" href="http://www.computerweekly.com/news/2240089572/Microsoft-UKs-national-technology-officer-moves-on">"Microsoft UK's national technology officer moves on"</a>. <i>Computer Weekly</i><span class="reference-accessdate">. Retrieved <span class="nowrap">17 December</span> 2016</span>.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Ajournal&amp;rft.genre=unknown&amp;rft.jtitle=Computer+Weekly&amp;rft.atitle=Microsoft+UK%27s+national+technology+officer+moves+on&amp;rft_id=http%3A%2F%2Fwww.computerweekly.com%2Fnews%2F2240089572%2FMicrosoft-UKs-national-technology-officer-moves-on&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
<li id="cite_note-51"><span class="mw-cite-backlink"><b><a href="#cite_ref-51">^</a></b></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite id="CITEREFFountain2021" class="citation web cs1">Fountain, Nigel (31 May 2021). <a rel="nofollow" class="external text" href="https://www.theguardian.com/science/2021/may/30/john-hodge-obituary">"John Hodge obituary"</a>. <i>The Guardian</i><span class="reference-accessdate">. Retrieved <span class="nowrap">11 October</span> 2021</span>.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Ajournal&amp;rft.genre=unknown&amp;rft.jtitle=The+Guardian&amp;rft.atitle=John+Hodge+obituary&amp;rft.date=2021-05-31&amp;rft.aulast=Fountain&amp;rft.aufirst=Nigel&amp;rft_id=https%3A%2F%2Fwww.theguardian.com%2Fscience%2F2021%2Fmay%2F30%2Fjohn-hodge-obituary&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
<li id="cite_note-52"><span class="mw-cite-backlink"><b><a href="#cite_ref-52">^</a></b></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite class="citation web cs1"><a rel="nofollow" class="external text" href="https://web.archive.org/web/20120725084803/http://www.tuc.org.uk/the_tuc/about_bbarber.cfm">"Biographical details: Brendan Barber"</a>. Trades Union Congress. Archived from <a rel="nofollow" class="external text" href="http://www.tuc.org.uk/the_tuc/about_bbarber.cfm">the original</a> on 25 July 2012<span class="reference-accessdate">. Retrieved <span class="nowrap">24 September</span> 2012</span>.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Abook&amp;rft.genre=unknown&amp;rft.btitle=Biographical+details%3A+Brendan+Barber&amp;rft.pub=Trades+Union+Congress&amp;rft_id=http%3A%2F%2Fwww.tuc.org.uk%2Fthe_tuc%2Fabout_bbarber.cfm&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
<li id="cite_note-PTC-53"><span class="mw-cite-backlink">^ <a href="#cite_ref-PTC_53-0"><sup><i><b>a</b></i></sup></a> <a href="#cite_ref-PTC_53-1"><sup><i><b>b</b></i></sup></a> <a href="#cite_ref-PTC_53-2"><sup><i><b>c</b></i></sup></a> <a href="#cite_ref-PTC_53-3"><sup><i><b>d</b></i></sup></a> <a href="#cite_ref-PTC_53-4"><sup><i><b>e</b></i></sup></a> <a href="#cite_ref-PTC_53-5"><sup><i><b>f</b></i></sup></a> <a href="#cite_ref-PTC_53-6"><sup><i><b>g</b></i></sup></a> <a href="#cite_ref-PTC_53-7"><sup><i><b>h</b></i></sup></a></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite class="citation web cs1"><a rel="nofollow" class="external text" href="https://archive.today/20130131200025/http://www.ptceducation.com/PROFILES/city.html">"City University- Institution Profiles"</a>. PTC. Archived from <a rel="nofollow" class="external text" href="http://www.ptceducation.com/PROFILES/city.html">the original</a> on 31 January 2013<span class="reference-accessdate">. Retrieved <span class="nowrap">30 November</span> 2012</span>.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Abook&amp;rft.genre=unknown&amp;rft.btitle=City+University-+Institution+Profiles&amp;rft.pub=PTC&amp;rft_id=http%3A%2F%2Fwww.ptceducation.com%2FPROFILES%2Fcity.html&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
<li id="cite_note-54"><span class="mw-cite-backlink"><b><a href="#cite_ref-54">^</a></b></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite id="CITEREFMinutaglio2019" class="citation web cs1">Minutaglio, Rose (18 March 2019). <a rel="nofollow" class="external text" href="https://www.townandcountrymag.com/society/a25846074/tiffany-trump-michael-boulos-relationship/">"Private Dining, Family Holiday Parties, and Clubbing: Inside Tiffany Trump's New Relationship"</a>. <i><a href="/wiki/Town_%26_Country_(magazine)" title="Town &amp; Country (magazine)">Town &amp; Country</a></i><span class="reference-accessdate">. Retrieved <span class="nowrap">4 August</span> 2020</span>.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Ajournal&amp;rft.genre=unknown&amp;rft.jtitle=Town+%26+Country&amp;rft.atitle=Private+Dining%2C+Family+Holiday+Parties%2C+and+Clubbing%3A+Inside+Tiffany+Trump%27s+New+Relationship&amp;rft.date=2019-03-18&amp;rft.aulast=Minutaglio&amp;rft.aufirst=Rose&amp;rft_id=https%3A%2F%2Fwww.townandcountrymag.com%2Fsociety%2Fa25846074%2Ftiffany-trump-michael-boulos-relationship%2F&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
<li id="cite_note-55"><span class="mw-cite-backlink"><b><a href="#cite_ref-55">^</a></b></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite class="citation web cs1"><a rel="nofollow" class="external text" href="http://www.dowjones.com/team/william-lewis-3/">"Dow Jones"</a>. <i>DowJones.com</i><span class="reference-accessdate">. Retrieved <span class="nowrap">20 November</span> 2017</span>.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Ajournal&amp;rft.genre=unknown&amp;rft.jtitle=DowJones.com&amp;rft.atitle=Dow+Jones&amp;rft_id=http%3A%2F%2Fwww.dowjones.com%2Fteam%2Fwilliam-lewis-3%2F&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
<li id="cite_note-Questex-56"><span class="mw-cite-backlink"><b><a href="#cite_ref-Questex_56-0">^</a></b></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite class="citation web cs1"><a rel="nofollow" class="external text" href="https://web.archive.org/web/20150208150928/http://www.berlinconference.com/index.php/speaker_profile/ian_livingstone">"Ian Livingstone"</a>. Questex Hospitality+Travel Group. Archived from <a rel="nofollow" class="external text" href="http://www.berlinconference.com/index.php/speaker_profile/ian_livingstone">the original</a> on 8 February 2015<span class="reference-accessdate">. Retrieved <span class="nowrap">8 February</span> 2015</span>.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Abook&amp;rft.genre=unknown&amp;rft.btitle=Ian+Livingstone&amp;rft.pub=Questex+Hospitality%2BTravel+Group&amp;rft_id=http%3A%2F%2Fwww.berlinconference.com%2Findex.php%2Fspeaker_profile%2Fian_livingstone&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
<li id="cite_note-57"><span class="mw-cite-backlink"><b><a href="#cite_ref-57">^</a></b></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite class="citation web cs1"><a rel="nofollow" class="external text" href="https://web.archive.org/web/20121026144423/http://www.igef.cuhk.edu.hk/index.php/en/people/professor-liu-mingkang/biography">"Professor Liu"</a>. Archived from <a rel="nofollow" class="external text" href="http://www.igef.cuhk.edu.hk/index.php/en/people/professor-liu-mingkang/biography">the original</a> on 26 October 2012<span class="reference-accessdate">. Retrieved <span class="nowrap">16 October</span> 2016</span>.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Abook&amp;rft.genre=unknown&amp;rft.btitle=Professor+Liu&amp;rft_id=http%3A%2F%2Fwww.igef.cuhk.edu.hk%2Findex.php%2Fen%2Fpeople%2Fprofessor-liu-mingkang%2Fbiography&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
<li id="cite_note-58"><span class="mw-cite-backlink"><b><a href="#cite_ref-58">^</a></b></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite class="citation web cs1"><a rel="nofollow" class="external text" href="https://www.citystgeorges.ac.uk/about/schools/communication-creativity/journalism">"Leading alumni... in newspapers"</a>. City University website. 8 July 2022.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Abook&amp;rft.genre=unknown&amp;rft.btitle=Leading+alumni...+in+newspapers&amp;rft.pub=City+University+website&amp;rft.date=2022-07-08&amp;rft_id=https%3A%2F%2Fwww.citystgeorges.ac.uk%2Fabout%2Fschools%2Fcommunication-creativity%2Fjournalism&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
<li id="cite_note-:02-59"><span class="mw-cite-backlink"><b><a href="#cite_ref-:02_59-0">^</a></b></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite id="CITEREFAnchunda2021" class="citation web cs1">Anchunda, Benly (13 August 2021). <a rel="nofollow" class="external text" href="https://www.crtv.cm/2021/08/2021-afcon-draw-your-co-host-multiple-award-winning-mimi-fawaz/">"2021 AFCON Draw: Your Co-host, Multiple Award Winning Mimi Fawaz"</a>. <i>Cameroon Radio Television</i>. <a rel="nofollow" class="external text" href="https://web.archive.org/web/20210813153628/https://www.crtv.cm/2021/08/2021-afcon-draw-your-co-host-multiple-award-winning-mimi-fawaz/">Archived</a> from the original on 13 August 2021<span class="reference-accessdate">. Retrieved <span class="nowrap">7 November</span> 2021</span>.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Ajournal&amp;rft.genre=unknown&amp;rft.jtitle=Cameroon+Radio+Television&amp;rft.atitle=2021+AFCON+Draw%3A+Your+Co-host%2C+Multiple+Award+Winning+Mimi+Fawaz&amp;rft.date=2021-08-13&amp;rft.aulast=Anchunda&amp;rft.aufirst=Benly&amp;rft_id=https%3A%2F%2Fwww.crtv.cm%2F2021%2F08%2F2021-afcon-draw-your-co-host-multiple-award-winning-mimi-fawaz%2F&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
<li id="cite_note-60"><span class="mw-cite-backlink"><b><a href="#cite_ref-60">^</a></b></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite class="citation web cs1"><a rel="nofollow" class="external text" href="https://www.tvnewsroom.co.uk/biography-images/lucrezia-millarini-10670/">"Lucrezia Millarini - biography and images"</a>. TV Newsroom. 8 June 2011<span class="reference-accessdate">. Retrieved <span class="nowrap">30 January</span> 2023</span>.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Abook&amp;rft.genre=unknown&amp;rft.btitle=Lucrezia+Millarini+-+biography+and+images&amp;rft.pub=TV+Newsroom&amp;rft.date=2011-06-08&amp;rft_id=https%3A%2F%2Fwww.tvnewsroom.co.uk%2Fbiography-images%2Flucrezia-millarini-10670%2F&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
<li id="cite_note-61"><span class="mw-cite-backlink"><b><a href="#cite_ref-61">^</a></b></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite class="citation web cs1"><a rel="nofollow" class="external text" href="https://www.citystgeorges.ac.uk/about/schools/communication-creativity/journalism">"Leading alumni in online and digital"</a>. <i>City, University of London</i><span class="reference-accessdate">. Retrieved <span class="nowrap">20 October</span> 2019</span>.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Ajournal&amp;rft.genre=unknown&amp;rft.jtitle=City%2C+University+of+London&amp;rft.atitle=Leading+alumni+in+online+and+digital&amp;rft_id=https%3A%2F%2Fwww.citystgeorges.ac.uk%2Fabout%2Fschools%2Fcommunication-creativity%2Fjournalism&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
<li id="cite_note-62"><span class="mw-cite-backlink"><b><a href="#cite_ref-62">^</a></b></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite id="CITEREFWintle2021" class="citation news cs1">Wintle, Angela (5 December 2021). <span class="id-lock-subscription" title="Paid subscription required"><a rel="nofollow" class="external text" href="https://www.telegraph.co.uk/money/fame-fortune/josh-widdicombe-basically-like-young-alan-sugar/">"Josh Widdicombe: 'I was basically like a young Alan Sugar'<span class="cs1-kern-right"></span>"</a></span>. <i><a href="/wiki/The_Daily_Telegraph" title="The Daily Telegraph">The Daily Telegraph</a></i>. <a href="/wiki/London" title="London">London</a>, England. <a rel="nofollow" class="external text" href="https://ghostarchive.org/archive/20220112/https://www.telegraph.co.uk/money/fame-fortune/josh-widdicombe-basically-like-young-alan-sugar/">Archived</a> from the original on 12 January 2022<span class="reference-accessdate">. Retrieved <span class="nowrap">5 December</span> 2021</span>.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Ajournal&amp;rft.genre=article&amp;rft.jtitle=The+Daily+Telegraph&amp;rft.atitle=Josh+Widdicombe%3A+%27I+was+basically+like+a+young+Alan+Sugar%27&amp;rft.date=2021-12-05&amp;rft.aulast=Wintle&amp;rft.aufirst=Angela&amp;rft_id=https%3A%2F%2Fwww.telegraph.co.uk%2Fmoney%2Ffame-fortune%2Fjosh-widdicombe-basically-like-young-alan-sugar%2F&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
<li id="cite_note-63"><span class="mw-cite-backlink"><b><a href="#cite_ref-63">^</a></b></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite class="citation web cs1"><a rel="nofollow" class="external text" href="http://www.ukgameshows.com/ukgs/Masterchef_Goes_Large">"Masterchef Goes Large - UKGameshows"</a>. <i>www.ukgameshows.com</i>.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Ajournal&amp;rft.genre=unknown&amp;rft.jtitle=www.ukgameshows.com&amp;rft.atitle=Masterchef+Goes+Large+-+UKGameshows&amp;rft_id=http%3A%2F%2Fwww.ukgameshows.com%2Fukgs%2FMasterchef_Goes_Large&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
<li id="cite_note-64"><span class="mw-cite-backlink"><b><a href="#cite_ref-64">^</a></b></span> <span class="reference-text"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1238218222" /><cite class="citation web cs1"><a rel="nofollow" class="external text" href="http://www.cosmopolitan.com/uk/entertainment/news/a43018/things-you-didnt-know-about-masterchef/">"MasterChef Studio"</a>. 29 April 2016.</cite><span title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Abook&amp;rft.genre=unknown&amp;rft.btitle=MasterChef+Studio&amp;rft.date=2016-04-29&amp;rft_id=http%3A%2F%2Fwww.cosmopolitan.com%2Fuk%2Fentertainment%2Fnews%2Fa43018%2Fthings-you-didnt-know-about-masterchef%2F&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ACity%2C+University+of+London" class="Z3988"></span></span>
</li>
</ol></div></div>
<div class="mw-heading mw-heading2"><h2 id="External_links">External links</h2><span class="mw-editsection"><span class="mw-editsection-bracket">[</span><a href="/w/index.php?title=City,_University_of_London&amp;action=edit&amp;section=35" title="Edit section: External links"><span>edit</span></a><span class="mw-editsection-bracket">]</span></span></div>
<style data-mw-deduplicate="TemplateStyles:r1290876196">.mw-parser-output .side-box{margin:4px 0;box-sizing:border-box;border:1px solid #aaa;font-size:88%;line-height:1.25em;background-color:var(--background-color-interactive-subtle,#f8f9fa);display:flow-root}.mw-parser-output .infobox .side-box{font-size:100%}.mw-parser-output .side-box-abovebelow,.mw-parser-output .side-box-text{padding:0.25em 0.9em}.mw-parser-output .side-box-image{padding:2px 0 2px 0.9em;text-align:center}.mw-parser-output .side-box-imageright{padding:2px 0.9em 2px 0;text-align:center}@media(min-width:500px){.mw-parser-output .side-box-flex{display:flex;align-items:center}.mw-parser-output .side-box-text{flex:1;min-width:0}}@media(min-width:720px){.mw-parser-output .side-box{width:238px}.mw-parser-output .side-box-right{clear:right;float:right;margin-left:1em}.mw-parser-output .side-box-left{margin-right:1em}}</style><style data-mw-deduplicate="TemplateStyles:r1237033735">@media print{body.ns-0 .mw-parser-output .sistersitebox{display:none!important}}@media screen{html.skin-theme-clientpref-night .mw-parser-output .sistersitebox img[src*="Wiktionary-logo-en-v2.svg"]{background-color:white}}@media screen and (prefers-color-scheme:dark){html.skin-theme-clientpref-os .mw-parser-output .sistersitebox img[src*="Wiktionary-logo-en-v2.svg"]{background-color:white}}</style><div class="side-box side-box-right plainlinks sistersitebox"><style data-mw-deduplicate="TemplateStyles:r1126788409">.mw-parser-output .plainlist ol,.mw-parser-output .plainlist ul{line-height:inherit;list-style:none;margin:0;padding:0}.mw-parser-output .plainlist ol li,.mw-parser-output .plainlist ul li{margin-bottom:0}</style>
<div class="side-box-flex">
<div class="side-box-image"><span class="noviewer" typeof="mw:File"><a href="/wiki/File:Commons-logo.svg" class="mw-file-description"><img alt="" src="//upload.wikimedia.org/wikipedia/en/thumb/4/4a/Commons-logo.svg/40px-Commons-logo.svg.png" decoding="async" width="30" height="40" class="mw-file-element" srcset="//upload.wikimedia.org/wikipedia/en/thumb/4/4a/Commons-logo.svg/60px-Commons-logo.svg.png 1.5x" data-file-width="1024" data-file-height="1376" /></a></span></div>
<div class="side-box-text plainlist">Wikimedia Commons has media related to <span style="font-weight: bold; font-style: italic;"><a href="https://commons.wikimedia.org/wiki/Category:City,_University_of_London" class="extiw" title="commons:Category:City, University of London">City, University of London</a></span>.</div></div>
</div>
<ul><li><a rel="nofollow" class="external text" href="http://www.citystgeorges.ac.uk/">City St George's, University of London</a></li>
<li><a rel="nofollow" class="external text" href="https://www.csgsu.co.uk/">City St George's, University of London, Students Union</a> <a rel="nofollow" class="external text" href="https://web.archive.org/web/20171021121607/http://www.culsu.co.uk/">Archived</a> 21 October 2017 at the <a href="/wiki/Wayback_Machine" title="Wayback Machine">Wayback Machine</a></li>
<li><a rel="nofollow" class="external text" href="https://web.archive.org/web/20100914105055/http://www.shl.lon.ac.uk/specialcollections/archives/studentrecords.shtml">Lists of Northampton Polytechnic Institute students</a></li>
<li><a rel="nofollow" class="external text" href="https://web.archive.org/web/20130510143909/http://www.senatehouselibrary.ac.uk/our-collections/historic-collections/archives-manuscripts/university-of-london-military-service-1914-1945/">List of Northampton Polytechnic Institute military personnel, 1914–1918</a></li></ul>
<div class="navbox-styles"><style data-mw-deduplicate="TemplateStyles:r1129693374">.mw-parser-output .hlist dl,.mw-parser-output .hlist ol,.mw-parser-output .hlist ul{margin:0;padding:0}.mw-parser-output .hlist dd,.mw-parser-output .hlist dt,.mw-parser-output .hlist li{margin:0;display:inline}.mw-parser-output .hlist.inline,.mw-parser-output .hlist.inline dl,.mw-parser-output .hlist.inline ol,.mw-parser-output .hlist.inline ul,.mw-parser-output .hlist dl dl,.mw-parser-output .hlist dl ol,.mw-parser-output .hlist dl ul,.mw-parser-output .hlist ol dl,.mw-parser-output .hlist ol ol,.mw-parser-output .hlist ol ul,.mw-parser-output .hlist ul dl,.mw-parser-output .hlist ul ol,.mw-parser-output .hlist ul ul{display:inline}.mw-parser-output .hlist .mw-empty-li{display:none}.mw-parser-output .hlist dt::after{content:": "}.mw-parser-output .hlist dd::after,.mw-parser-output .hlist li::after{content:" · ";font-weight:bold}.mw-parser-output .hlist dd:last-child::after,.mw-parser-output .hlist dt:last-child::after,.mw-parser-output .hlist li:last-child::after{content:none}.mw-parser-output .hlist dd dd:first-child::before,.mw-parser-output .hlist dd dt:first-child::before,.mw-parser-output .hlist dd li:first-child::before,.mw-parser-output .hlist dt dd:first-child::before,.mw-parser-output .hlist dt dt:first-child::before,.mw-parser-output .hlist dt li:first-child::before,.mw-parser-output .hlist li dd:first-child::before,.mw-parser-output .hlist li dt:first-child::before,.mw-parser-output .hlist li li:first-child::before{content:" (";font-weight:normal}.mw-parser-output .hlist dd dd:last-child::after,.mw-parser-output .hlist dd dt:last-child::after,.mw-parser-output .hlist dd li:last-child::after,.mw-parser-output .hlist dt dd:last-child::after,.mw-parser-output .hlist dt dt:last-child::after,.mw-parser-output .hlist dt li:last-child::after,.mw-parser-output .hlist li dd:last-child::after,.mw-parser-output .hlist li dt:last-child::after,.mw-parser-output .hlist li li:last-child::after{content:")";font-weight:normal}.mw-parser-output .hlist ol{counter-reset:listitem}.mw-parser-output .hlist ol>li{counter-increment:listitem}.mw-parser-output .hlist ol>li::before{content:" "counter(listitem)"\a0 "}.mw-parser-output .hlist dd ol>li:first-child::before,.mw-parser-output .hlist dt ol>li:first-child::before,.mw-parser-output .hlist li ol>li:first-child::before{content:" ("counter(listitem)"\a0 "}</style><style data-mw-deduplicate="TemplateStyles:r1236075235">.mw-parser-output .navbox{box-sizing:border-box;border:1px solid #a2a9b1;width:100%;clear:both;font-size:88%;text-align:center;padding:1px;margin:1em auto 0}.mw-parser-output .navbox .navbox{margin-top:0}.mw-parser-output .navbox+.navbox,.mw-parser-output .navbox+.navbox-styles+.navbox{margin-top:-1px}.mw-parser-output .navbox-inner,.mw-parser-output .navbox-subgroup{width:100%}.mw-parser-output .navbox-group,.mw-parser-output .navbox-title,.mw-parser-output .navbox-abovebelow{padding:0.25em 1em;line-height:1.5em;text-align:center}.mw-parser-output .navbox-group{white-space:nowrap;text-align:right}.mw-parser-output .navbox,.mw-parser-output .navbox-subgroup{background-color:#fdfdfd}.mw-parser-output .navbox-list{line-height:1.5em;border-color:#fdfdfd}.mw-parser-output .navbox-list-with-group{text-align:left;border-left-width:2px;border-left-style:solid}.mw-parser-output tr+tr>.navbox-abovebelow,.mw-parser-output tr+tr>.navbox-group,.mw-parser-output tr+tr>.navbox-image,.mw-parser-output tr+tr>.navbox-list{border-top:2px solid #fdfdfd}.mw-parser-output .navbox-title{background-color:#ccf}.mw-parser-output .navbox-abovebelow,.mw-parser-output .navbox-group,.mw-parser-output .navbox-subgroup .navbox-title{background-color:#ddf}.mw-parser-output .navbox-subgroup .navbox-group,.mw-parser-output .navbox-subgroup .navbox-abovebelow{background-color:#e6e6ff}.mw-parser-output .navbox-even{background-color:#f7f7f7}.mw-parser-output .navbox-odd{background-color:transparent}.mw-parser-output .navbox .hlist td dl,.mw-parser-output .navbox .hlist td ol,.mw-parser-output .navbox .hlist td ul,.mw-parser-output .navbox td.hlist dl,.mw-parser-output .navbox td.hlist ol,.mw-parser-output .navbox td.hlist ul{padding:0.125em 0}.mw-parser-output .navbox .navbar{display:block;font-size:100%}.mw-parser-output .navbox-title .navbar{float:left;text-align:left;margin-right:0.5em}body.skin--responsive .mw-parser-output .navbox-image img{max-width:none!important}@media print{body.ns-0 .mw-parser-output .navbox{display:none!important}}</style><style data-mw-deduplicate="TemplateStyles:r1239334494">@media screen{html.skin-theme-clientpref-night .mw-parser-output div:not(.notheme)>.tmp-color,html.skin-theme-clientpref-night .mw-parser-output p>.tmp-color,html.skin-theme-clientpref-night .mw-parser-output table:not(.notheme) .tmp-color{color:inherit!important}}@media screen and (prefers-color-scheme:dark){html.skin-theme-clientpref-os .mw-parser-output div:not(.notheme)>.tmp-color,html.skin-theme-clientpref-os .mw-parser-output p>.tmp-color,html.skin-theme-clientpref-os .mw-parser-output table:not(.notheme) .tmp-color{color:inherit!important}}</style><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1239334494" /></div><div role="navigation" class="navbox" aria-labelledby="City,_University_of_London342" style="padding:3px"><table class="nowraplinks hlist mw-collapsible expanded navbox-inner" style="border-spacing:0;background:transparent;color:inherit"><tbody><tr><th scope="col" class="navbox-title" colspan="2" style="background:#C00F22; color:#FFFFFF;box-shadow: inset 2px 2px 0 black, inset -2px -2px 0 black;"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1129693374" /><style data-mw-deduplicate="TemplateStyles:r1239400231">.mw-parser-output .navbar{display:inline;font-size:88%;font-weight:normal}.mw-parser-output .navbar-collapse{float:left;text-align:left}.mw-parser-output .navbar-boxtext{word-spacing:0}.mw-parser-output .navbar ul{display:inline-block;white-space:nowrap;line-height:inherit}.mw-parser-output .navbar-brackets::before{margin-right:-0.125em;content:"[ "}.mw-parser-output .navbar-brackets::after{margin-left:-0.125em;content:" ]"}.mw-parser-output .navbar li{word-spacing:-0.125em}.mw-parser-output .navbar a>span,.mw-parser-output .navbar a>abbr{text-decoration:inherit}.mw-parser-output .navbar-mini abbr{font-variant:small-caps;border-bottom:none;text-decoration:none;cursor:inherit}.mw-parser-output .navbar-ct-full{font-size:114%;margin:0 7em}.mw-parser-output .navbar-ct-mini{font-size:114%;margin:0 4em}html.skin-theme-clientpref-night .mw-parser-output .navbar li a abbr{color:var(--color-base)!important}@media(prefers-color-scheme:dark){html.skin-theme-clientpref-os .mw-parser-output .navbar li a abbr{color:var(--color-base)!important}}@media print{.mw-parser-output .navbar{display:none!important}}</style><div class="navbar plainlinks hlist navbar-mini"><ul><li class="nv-view"><a href="/wiki/Template:City,_University_of_London" title="Template:City, University of London"><abbr title="View this template" style="color:#FFFFFF">v</abbr></a></li><li class="nv-talk"><a href="/wiki/Template_talk:City,_University_of_London" title="Template talk:City, University of London"><abbr title="Discuss this template" style="color:#FFFFFF">t</abbr></a></li><li class="nv-edit"><a href="/wiki/Special:EditPage/Template:City,_University_of_London" title="Special:EditPage/Template:City, University of London"><abbr title="Edit this template" style="color:#FFFFFF">e</abbr></a></li></ul></div><div id="City,_University_of_London342" style="font-size:114%;margin:0 4em"><a class="mw-selflink selflink"><span class="tmp-color" style="color:white">City, University of London</span></a></div></th></tr><tr><td class="navbox-abovebelow" colspan="2" style="background:#C00F22; color:#FFFFFF;box-shadow: inset 2px 2px 0 black, inset -2px -2px 0 black;"><div><a href="/wiki/University_of_London" title="University of London"><span class="tmp-color" style="color:#FFFFFF">University of London</span></a></div></td></tr><tr><th scope="row" class="navbox-group" style="background:#C00F22; color:#FFFFFF;box-shadow: inset 2px 2px 0 black, inset -2px -2px 0 black;;width:1%">Subdivisions and<br />publications</th><td class="navbox-list-with-group navbox-list navbox-odd" style="width:100%;padding:0"><div style="padding:0 0.25em">
<ul><li><a href="/wiki/Bayes_Business_School" title="Bayes Business School">Bayes Business School</a></li>
<li><a href="/wiki/City_Law_School" title="City Law School">City Law School</a>
<ul><li><a href="/wiki/Inns_of_Court_School_of_Law" class="mw-redirect" title="Inns of Court School of Law">Inns of Court School of Law</a></li></ul></li>
<li>School of Arts
<ul><li><a href="/wiki/Department_of_Journalism,_City_University" title="Department of Journalism, City University">Journalism Department</a></li></ul></li>
<li><a href="/wiki/School_of_Health_Sciences,_City_University_London" class="mw-redirect" title="School of Health Sciences, City University London">School of Community and Health Sciences</a>
<ul><li><a href="/wiki/St_Bartholomew_School_of_Nursing_%26_Midwifery" class="mw-redirect" title="St Bartholomew School of Nursing &amp; Midwifery">St Bartholomew School of Nursing &amp; Midwifery</a></li></ul></li>
<li>School of Engineering and Mathematical Sciences</li>
<li>School of Informatics
<ul><li><a href="/wiki/Centre_for_Software_Reliability" title="Centre for Software Reliability">Centre for Software Reliability</a></li></ul></li>
<li>School of Social Sciences</li></ul>
</div></td></tr><tr><th scope="row" class="navbox-group" style="background:#C00F22; color:#FFFFFF;box-shadow: inset 2px 2px 0 black, inset -2px -2px 0 black;;width:1%">University</th><td class="navbox-list-with-group navbox-list navbox-odd" style="width:100%;padding:0"><div style="padding:0 0.25em"></div><table class="nowraplinks navbox-subgroup" style="border-spacing:0"><tbody><tr><th scope="row" class="navbox-group" style="background:#C00F22; color:#FFFFFF;box-shadow: inset 2px 2px 0 black, inset -2px -2px 0 black;;width:1%"><a class="mw-selflink-fragment" href="#Campus"><span style="color:white">Campus</span></a></th><td class="navbox-list-with-group navbox-list navbox-even" style="width:100%;padding:0"><div style="padding:0 0.25em">
<ul><li><a href="/wiki/Northampton_Square" title="Northampton Square">Northampton Square</a></li>
<li><a href="/wiki/Old_Street" title="Old Street">Old Street</a></li>
<li><a href="/wiki/A1_road_(London)#St_John_Street_(historic)" class="mw-redirect" title="A1 road (London)">St John Street</a></li></ul>
</div></td></tr><tr><th scope="row" class="navbox-group" style="background:#C00F22; color:#FFFFFF;box-shadow: inset 2px 2px 0 black, inset -2px -2px 0 black;;width:1%"><a href="/wiki/Category:People_associated_with_City,_University_of_London" title="Category:People associated with City, University of London"><span style="color:white">People</span></a></th><td class="navbox-list-with-group navbox-list navbox-odd" style="width:100%;padding:0"><div style="padding:0 0.25em">
<ul><li><a href="/wiki/Category:Academics_of_City,_University_of_London" title="Category:Academics of City, University of London">Academics</a></li>
<li><a href="/wiki/City_University,_London#Notable_academics_and_alumni" class="mw-redirect" title="City University, London">Alumni</a></li>
<li><i>Rector:</i> <a href="/wiki/Lord_Mayor_of_London" title="Lord Mayor of London">Lord Mayor of London</a> (ex officio)</li>
<li><i>President:</i> <a href="/wiki/Anthony_Finkelstein" title="Anthony Finkelstein">Anthony Finkelstein</a></li></ul>
</div></td></tr><tr><th scope="row" class="navbox-group" style="background:#C00F22; color:#FFFFFF;box-shadow: inset 2px 2px 0 black, inset -2px -2px 0 black;;width:1%">Other</th><td class="navbox-list-with-group navbox-list navbox-even" style="width:100%;padding:0"><div style="padding:0 0.25em">
<ul><li><a href="/wiki/Contrarian_Prize#Lectures" title="Contrarian Prize">Contrarian Lectures</a></li>
<li><a href="/wiki/City_University_London#History" class="mw-redirect" title="City University London">History</a></li>
<li><a href="/wiki/Mais_Lecture" title="Mais Lecture">Mais Lectures</a></li>
<li><a href="/wiki/City_University_London#Students&#39;_Union" class="mw-redirect" title="City University London">Students' Union</a></li></ul>
</div></td></tr></tbody></table><div></div></td></tr><tr><th scope="row" class="navbox-group" style="background:#C00F22; color:#FFFFFF;box-shadow: inset 2px 2px 0 black, inset -2px -2px 0 black;;width:1%">Affiliates</th><td class="navbox-list-with-group navbox-list navbox-odd" style="width:100%;padding:0"><div style="padding:0 0.25em">
<ul><li><a href="/wiki/Information_Centre_about_Asylum_and_Refugees" title="Information Centre about Asylum and Refugees">Information Centre about Asylum and Refugees</a></li>
<li><a href="/wiki/London_Centre_for_Arts_and_Cultural_Exchange" title="London Centre for Arts and Cultural Exchange">London Centre for Arts and Cultural Exchange</a></li>
<li><a href="/wiki/City_University_London#Partnerships_and_collaborations" class="mw-redirect" title="City University London">Partnerships and collaborations</a></li>
<li><a href="/wiki/Universities_UK" title="Universities UK">Universities UK</a></li></ul>
</div></td></tr></tbody></table></div>
<div class="navbox-styles"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1236075235" /></div><div role="navigation" class="navbox" aria-labelledby="Links_to_related_articles92" style="padding:3px"><table class="nowraplinks mw-collapsible mw-collapsed navbox-inner" style="border-spacing:0;background:transparent;color:inherit"><tbody><tr><th scope="col" class="navbox-title" colspan="2" style="background:#e8e8ff;"><div id="Links_to_related_articles92" style="font-size:114%;margin:0 4em">Links to related articles</div></th></tr><tr><td colspan="2" class="navbox-list navbox-odd" style="width:100%;padding:0;font-size:114%"><div style="padding:0px">
<div class="navbox-styles"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1129693374" /><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1236075235" /><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1239334494" /><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1239334494" /><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1239334494" /><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1239334494" /><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1239334494" /><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1239334494" /></div><div role="navigation" class="navbox" aria-labelledby="University_of_London1098" style="padding:3px"><table class="nowraplinks hlist mw-collapsible autocollapse navbox-inner" style="border-spacing:0;background:transparent;color:inherit"><tbody><tr><th scope="col" class="navbox-title" colspan="2" style="background: #003F87; color: white; box-shadow: inset 2px 2px 0 #ED2939, inset -2px -2px 0 #ED2939;"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1129693374" /><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1239400231" /><div class="navbar plainlinks hlist navbar-mini"><ul><li class="nv-view"><a href="/wiki/Template:University_of_London" title="Template:University of London"><abbr title="View this template" style="color: white">v</abbr></a></li><li class="nv-talk"><a href="/wiki/Template_talk:University_of_London" title="Template talk:University of London"><abbr title="Discuss this template" style="color: white">t</abbr></a></li><li class="nv-edit"><a href="/wiki/Special:EditPage/Template:University_of_London" title="Special:EditPage/Template:University of London"><abbr title="Edit this template" style="color: white">e</abbr></a></li></ul></div><div id="University_of_London1098" style="font-size:114%;margin:0 4em"><a href="/wiki/University_of_London" title="University of London"><span class="tmp-color" style="color:#E5EEF8">University of London</span></a></div></th></tr><tr><th scope="row" class="navbox-group" style="background: #003F87; color: white; box-shadow: inset 2px 2px 0 #ED2939, inset -2px -2px 0 #ED2939;;width:1%;background: #ED2939; color: white"><a href="/wiki/University_of_London#Administrative_structure" title="University of London"><span class="tmp-color" style="color:white">Institutions</span></a></th><td class="navbox-list-with-group navbox-list navbox-odd" style="width:100%;padding:0"><div style="padding:0 0.25em">
<ul><li><a href="/wiki/Birkbeck,_University_of_London" title="Birkbeck, University of London">Birkbeck</a></li>
<li><a href="/wiki/Brunel_University_of_London" title="Brunel University of London">Brunel</a></li>
<li><a href="/wiki/City_St_George%27s,_University_of_London" title="City St George&#39;s, University of London">City St George's</a></li>
<li><a href="/wiki/Courtauld_Institute_of_Art" title="Courtauld Institute of Art">Courtauld Institute of Art</a></li>
<li><a href="/wiki/Goldsmiths,_University_of_London" title="Goldsmiths, University of London">Goldsmiths</a></li>
<li><a href="/wiki/Institute_of_Cancer_Research" title="Institute of Cancer Research">Institute of Cancer Research</a></li>
<li><a href="/wiki/King%27s_College_London" title="King&#39;s College London">King's College London</a></li>
<li><a href="/wiki/London_Business_School" title="London Business School">London Business School</a></li>
<li><a href="/wiki/London_School_of_Economics" title="London School of Economics">London School of Economics</a></li>
<li><a href="/wiki/London_School_of_Hygiene_%26_Tropical_Medicine" title="London School of Hygiene &amp; Tropical Medicine">London School of Hygiene &amp; Tropical Medicine</a></li>
<li><a href="/wiki/Queen_Mary_University_of_London" title="Queen Mary University of London">Queen Mary</a></li>
<li><a href="/wiki/Royal_Academy_of_Music" title="Royal Academy of Music">Royal Academy of Music</a></li>
<li><a href="/wiki/Royal_Central_School_of_Speech_and_Drama" title="Royal Central School of Speech and Drama">Royal Central School of Speech and Drama</a></li>
<li><a href="/wiki/Royal_Holloway,_University_of_London" title="Royal Holloway, University of London">Royal Holloway</a></li>
<li><a href="/wiki/Royal_Veterinary_College" title="Royal Veterinary College">Royal Veterinary College</a></li>
<li><a href="/wiki/SOAS_University_of_London" title="SOAS University of London">School of Oriental and African Studies</a></li>
<li><a href="/wiki/University_College_London" title="University College London">University College London</a></li></ul>
</div></td></tr><tr><th scope="row" class="navbox-group" style="background: #003F87; color: white; box-shadow: inset 2px 2px 0 #ED2939, inset -2px -2px 0 #ED2939;;width:1%;background: #ED2939; color: white">Central bodies<br />and programmes</th><td class="navbox-list-with-group navbox-list navbox-even" style="width:100%;padding:0"><div style="padding:0 0.25em">
<ul><li><a href="/wiki/Senate_House_Libraries" title="Senate House Libraries">Senate House Libraries</a></li>
<li><a href="/wiki/School_of_Advanced_Study" title="School of Advanced Study">School of Advanced Study</a>
<ul><li><a href="/wiki/Institute_of_Advanced_Legal_Studies" title="Institute of Advanced Legal Studies">Institute of Advanced Legal Studies</a></li>
<li><a href="/wiki/Institute_of_Classical_Studies" title="Institute of Classical Studies">Institute of Classical Studies</a></li>
<li><a href="/wiki/Institute_of_Commonwealth_Studies" title="Institute of Commonwealth Studies">Institute of Commonwealth Studies</a></li>
<li><a href="/wiki/Institute_of_English_Studies" title="Institute of English Studies">Institute of English Studies</a></li>
<li><a href="/wiki/Institute_of_Historical_Research" title="Institute of Historical Research">Institute of Historical Research</a>
<ul><li><a href="/wiki/Centre_for_Metropolitan_History" title="Centre for Metropolitan History">Centre for Metropolitan History</a></li></ul></li>
<li><a href="/wiki/Institute_of_Latin_American_Studies" title="Institute of Latin American Studies">Institute of Latin American Studies</a></li>
<li><a href="/wiki/Institute_of_Modern_Languages_Research" title="Institute of Modern Languages Research">Institute of Modern Languages Research</a></li>
<li><a href="/wiki/Institute_of_Philosophy,_University_of_London" title="Institute of Philosophy, University of London">Institute of Philosophy</a></li>
<li><a href="/wiki/Warburg_Institute" title="Warburg Institute">Warburg Institute</a></li></ul></li>
<li><a href="/wiki/University_of_London_Institute_in_Paris" title="University of London Institute in Paris">University of London Institute in Paris</a></li>
<li><a href="/wiki/University_of_London_Worldwide" title="University of London Worldwide">University of London Worldwide</a></li></ul>
</div></td></tr><tr><th scope="row" class="navbox-group" style="background: #003F87; color: white; box-shadow: inset 2px 2px 0 #ED2939, inset -2px -2px 0 #ED2939;;width:1%;background: #ED2939; color: white"><a href="/wiki/Category:People_associated_with_the_University_of_London" title="Category:People associated with the University of London"><span class="tmp-color" style="color:white">People</span></a></th><td class="navbox-list-with-group navbox-list navbox-odd" style="width:100%;padding:0"><div style="padding:0 0.25em">
<ul><li><i><b><a href="/wiki/List_of_Chancellors_of_the_University_of_London" class="mw-redirect" title="List of Chancellors of the University of London">Chancellor</a>:</b></i> <a href="/wiki/Anne,_Princess_Royal" title="Anne, Princess Royal">The Princess Royal</a></li>
<li><i><b><a href="/wiki/List_of_vice-chancellors_of_the_University_of_London" title="List of vice-chancellors of the University of London">Vice-Chancellor</a>:</b></i> <a href="/wiki/Wendy_Thomson" title="Wendy Thomson">Wendy Thomson</a></li>
<li><i><b><a href="/wiki/Visitor" title="Visitor">Visitor</a>:</b></i> <a href="/wiki/Lord_President_of_the_Council" title="Lord President of the Council">The Lord President of the Council</a> (<a href="/wiki/Lucy_Powell" title="Lucy Powell">Lucy Powell</a>)</li></ul>
<ul><li><a href="/wiki/Category:Academics_of_the_University_of_London" title="Category:Academics of the University of London">Academics</a></li>
<li><a href="/wiki/Category:Alumni_of_the_University_of_London" title="Category:Alumni of the University of London">Alumni</a></li>
<li><a href="/wiki/List_of_heads_of_member_institutions_of_the_University_of_London" title="List of heads of member institutions of the University of London">List of heads of colleges</a></li>
<li><a href="/wiki/List_of_University_of_London_people" class="mw-redirect" title="List of University of London people">List of University of London people</a></li></ul>
</div></td></tr><tr><th scope="row" class="navbox-group" style="background: #003F87; color: white; box-shadow: inset 2px 2px 0 #ED2939, inset -2px -2px 0 #ED2939;;width:1%;background: #ED2939; color: white">Places and<br />buildings</th><td class="navbox-list-with-group navbox-list navbox-even" style="width:100%;padding:0"><div style="padding:0 0.25em">
<ul><li><a href="/wiki/Bloomsbury" title="Bloomsbury">Bloomsbury</a></li>
<li><a href="/wiki/Gordon_Square" title="Gordon Square">Gordon Square</a></li>
<li><a href="/wiki/Category:University_of_London_intercollegiate_halls_of_residence" title="Category:University of London intercollegiate halls of residence">Halls of residence</a>
<ul><li><a href="/wiki/College_Hall,_London" title="College Hall, London">College Hall</a></li>
<li><a href="/wiki/Connaught_Hall,_London" title="Connaught Hall, London">Connaught Hall</a></li>
<li><a href="/wiki/International_Hall,_London" title="International Hall, London">International Hall</a></li>
<li><a href="/wiki/Nutford_House,_London" title="Nutford House, London">Nutford House</a></li></ul></li>
<li><a href="/wiki/Malet_Street" title="Malet Street">Malet Street</a></li>
<li><a href="/wiki/Russell_Square" title="Russell Square">Russell Square</a></li>
<li><a href="/wiki/Senate_House,_London" title="Senate House, London">Senate House</a></li>
<li><a href="/wiki/Tavistock_Square" title="Tavistock Square">Tavistock Square</a></li>
<li><a href="/wiki/Torrington_Square" title="Torrington Square">Torrington Square</a></li>
<li><a href="/wiki/Woburn_Square" title="Woburn Square">Woburn Square</a></li></ul>
</div></td></tr><tr><th scope="row" class="navbox-group" style="background: #003F87; color: white; box-shadow: inset 2px 2px 0 #ED2939, inset -2px -2px 0 #ED2939;;width:1%;background: #ED2939; color: white"><a href="/wiki/University_of_London#History" title="University of London"><span class="tmp-color" style="color:white">History</span></a></th><td class="navbox-list-with-group navbox-list navbox-odd" style="width:100%;padding:0"><div style="padding:0 0.25em"></div><table class="nowraplinks navbox-subgroup" style="border-spacing:0"><tbody><tr><th scope="row" class="navbox-group" style="width:1%;background:white; color: #003F87; box-shadow: inset 1px 1px 0 #ED2939, inset -1px -1px 0 #ED2939;"><a href="/wiki/Category:Former_colleges_of_the_University_of_London" title="Category:Former colleges of the University of London"><span class="tmp-color" style="color:#003F87">Institutions</span></a></th><td class="navbox-list-with-group navbox-list navbox-odd" style="width:100%;padding:0"><div style="padding:0 0.25em">
<ul><li><a href="/wiki/Bedford_College,_London" title="Bedford College, London">Bedford College</a></li>
<li><a href="/wiki/Chelsea_College_of_Science_and_Technology" title="Chelsea College of Science and Technology">Chelsea College of Science and Technology</a></li>
<li><a class="mw-selflink selflink">City</a></li>
<li><a href="/wiki/Heythrop_College,_University_of_London" title="Heythrop College, University of London">Heythrop College</a></li>
<li><a href="/wiki/Imperial_College_London" title="Imperial College London">Imperial College</a></li>
<li><a href="/wiki/UCL_Institute_of_Education" title="UCL Institute of Education">Institute of Education</a></li>
<li><a href="/wiki/Institute_of_Psychiatry,_Psychology_and_Neuroscience" title="Institute of Psychiatry, Psychology and Neuroscience">Institute of Psychiatry, Psychology and Neuroscience</a></li>
<li><a href="/wiki/Institute_for_the_Study_of_the_Americas" title="Institute for the Study of the Americas">Institute for the Study of the Americas</a></li>
<li><a href="/wiki/London_Consortium" title="London Consortium">London Consortium</a></li>
<li><a href="/wiki/New_College_London" title="New College London">New College</a></li>
<li><a href="/wiki/Queen_Elizabeth_College" title="Queen Elizabeth College">Queen Elizabeth College</a></li>
<li><a href="/wiki/Regent%27s_Park_College,_Oxford" title="Regent&#39;s Park College, Oxford">Regent's Park College</a></li>
<li><a href="/wiki/Royal_Postgraduate_Medical_School" title="Royal Postgraduate Medical School">Royal Postgraduate Medical School</a></li>
<li><a href="/wiki/St_George%27s,_University_of_London" title="St George&#39;s, University of London">St George's</a></li>
<li><a href="/wiki/UCL_School_of_Pharmacy" title="UCL School of Pharmacy">School of Pharmacy</a></li>
<li><a href="/wiki/UCL_School_of_Slavonic_and_East_European_Studies" title="UCL School of Slavonic and East European Studies">School of Slavonic and East European Studies</a></li>
<li><a href="/wiki/University_Marine_Biological_Station_Millport" title="University Marine Biological Station Millport">University Marine Biological Station Millport</a></li>
<li><a href="/wiki/St_Thomas%27s_Hospital_Medical_School" title="St Thomas&#39;s Hospital Medical School">St Thomas's Hospital Medical School</a></li>
<li><a href="/wiki/Westfield_College" title="Westfield College">Westfield College</a></li>
<li><a href="/wiki/Wye_College" title="Wye College">Wye College</a></li></ul>
</div></td></tr><tr><th scope="row" class="navbox-group" style="width:1%;background:white; color: #003F87; box-shadow: inset 1px 1px 0 #ED2939, inset -1px -1px 0 #ED2939;">Buildings</th><td class="navbox-list-with-group navbox-list navbox-even" style="width:100%;padding:0"><div style="padding:0 0.25em">
<ul><li><a href="/wiki/6_Burlington_Gardens" title="6 Burlington Gardens">6 Burlington Gardens</a></li>
<li><a href="/wiki/Church_of_Christ_the_King,_Bloomsbury" title="Church of Christ the King, Bloomsbury">Church of Christ the King</a></li>
<li><a href="/wiki/Category:University_of_London_intercollegiate_halls_of_residence" title="Category:University of London intercollegiate halls of residence">Halls of residence</a>
<ul><li><a href="/wiki/Commonwealth_Hall" title="Commonwealth Hall">Commonwealth Hall</a></li>
<li><a href="/wiki/Hughes_Parry_Hall,_London" title="Hughes Parry Hall, London">Hughes Parry Hall</a></li></ul></li></ul>
</div></td></tr><tr><td colspan="2" class="navbox-list navbox-odd" style="width:100%;padding:0"><div style="padding:0 0.25em">
<ul><li><a href="/wiki/London_University_(UK_Parliament_constituency)" title="London University (UK Parliament constituency)">Parliamentary constituency</a></li>
<li><a href="/wiki/General_Examination_for_Women" title="General Examination for Women">General Examination for Women</a></li>
<li><a href="/wiki/Privileged_bodies_of_the_United_Kingdom" title="Privileged bodies of the United Kingdom">Privileged bodies of the United Kingdom</a></li></ul>
</div></td></tr></tbody></table><div></div></td></tr><tr><th scope="row" class="navbox-group" style="background: #003F87; color: white; box-shadow: inset 2px 2px 0 #ED2939, inset -2px -2px 0 #ED2939;;width:1%;background: #ED2939; color: white">Other</th><td class="navbox-list-with-group navbox-list navbox-even" style="width:100%;padding:0"><div style="padding:0 0.25em">
<ul><li><a href="/wiki/Academic_dress_of_the_University_of_London" title="Academic dress of the University of London">Academic dress</a></li>
<li><a href="/wiki/The_Careers_Group,_University_of_London" title="The Careers Group, University of London">The Careers Group</a></li>
<li><i><a href="/wiki/London_Student" title="London Student">London Student</a></i></li>
<li><a href="/wiki/University_of_London_Boat_Club" title="University of London Boat Club">University of London Boat Club</a></li>
<li><a href="/wiki/University_of_London_Computer_Centre" title="University of London Computer Centre">University of London Computer Centre</a></li>
<li><a href="/wiki/University_of_London_Press" title="University of London Press">University of London Press</a></li>
<li><a href="/wiki/University_of_London_Union" title="University of London Union">University of London Union</a></li></ul>
</div></td></tr><tr><td class="navbox-abovebelow" colspan="2" style="background: #003F87; color: white; box-shadow: inset 2px 2px 0 #ED2939, inset -2px -2px 0 #ED2939;"><div>
<ul><li><span class="noviewer" typeof="mw:File"><span title="Category"><img alt="" src="//upload.wikimedia.org/wikipedia/en/thumb/9/96/Symbol_category_class.svg/20px-Symbol_category_class.svg.png" decoding="async" width="16" height="16" class="mw-file-element" srcset="//upload.wikimedia.org/wikipedia/en/thumb/9/96/Symbol_category_class.svg/40px-Symbol_category_class.svg.png 1.5x" data-file-width="180" data-file-height="185" /></span></span> <b><a href="/wiki/Category:University_of_London" title="Category:University of London"><span class="tmp-color" style="color:white">Category</span></a></b></li>
<li><span class="noviewer" typeof="mw:File"><span title="Commons page"><img alt="" src="//upload.wikimedia.org/wikipedia/en/thumb/4/4a/Commons-logo.svg/20px-Commons-logo.svg.png" decoding="async" width="12" height="16" class="mw-file-element" srcset="//upload.wikimedia.org/wikipedia/en/thumb/4/4a/Commons-logo.svg/40px-Commons-logo.svg.png 2x" data-file-width="1024" data-file-height="1376" /></span></span> <b><a href="https://commons.wikimedia.org/wiki/Category:University_of_London" class="extiw" title="commons:Category:University of London"><span class="tmp-color" style="color:white">Commons</span></a></b></li></ul>
</div></td></tr></tbody></table></div>
<div class="navbox-styles"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1129693374" /><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1236075235" /></div><div role="navigation" class="navbox" aria-labelledby="Universities_and_colleges_in_London353" style="padding:3px"><table class="nowraplinks mw-collapsible mw-collapsed navbox-inner" style="border-spacing:0;background:transparent;color:inherit"><tbody><tr><th scope="col" class="navbox-title" colspan="2"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1129693374" /><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1239400231" /><div class="navbar plainlinks hlist navbar-mini"><ul><li class="nv-view"><a href="/wiki/Template:Universities_and_colleges_in_London" title="Template:Universities and colleges in London"><abbr title="View this template">v</abbr></a></li><li class="nv-talk"><a href="/wiki/Template_talk:Universities_and_colleges_in_London" title="Template talk:Universities and colleges in London"><abbr title="Discuss this template">t</abbr></a></li><li class="nv-edit"><a href="/wiki/Special:EditPage/Template:Universities_and_colleges_in_London" title="Special:EditPage/Template:Universities and colleges in London"><abbr title="Edit this template">e</abbr></a></li></ul></div><div id="Universities_and_colleges_in_London353" style="font-size:114%;margin:0 4em">Universities and colleges in <a href="/wiki/London" title="London">London</a></div></th></tr><tr><td class="navbox-abovebelow" colspan="2"><div><a href="/wiki/Education_in_London" title="Education in London">Education in London</a></div></td></tr><tr><th scope="row" class="navbox-group" style="width:1%"><a href="/wiki/Category:Universities_and_colleges_in_London" title="Category:Universities and colleges in London">Universities</a></th><td class="navbox-list-with-group navbox-list navbox-odd hlist" style="width:100%;padding:0"><div style="padding:0 0.25em"></div><table class="nowraplinks navbox-subgroup" style="border-spacing:0"><tbody><tr><th scope="row" class="navbox-group" style="width:1%"><a href="/wiki/University_of_London" title="University of London">University <br /> of London</a></th><td class="navbox-list-with-group navbox-list navbox-odd" style="width:100%;padding:0"><div style="padding:0 0.25em">
<ul><li><a href="/wiki/Birkbeck,_University_of_London" title="Birkbeck, University of London">Birkbeck</a></li>
<li><a href="/wiki/Brunel_University_London" class="mw-redirect" title="Brunel University London">Brunel</a></li>
<li><a href="/wiki/City_St_George%27s,_University_of_London" title="City St George&#39;s, University of London">City St George's</a></li>
<li><a href="/wiki/Courtauld_Institute_of_Art" title="Courtauld Institute of Art">Courtauld Institute of Art</a></li>
<li><a href="/wiki/Goldsmiths,_University_of_London" title="Goldsmiths, University of London">Goldsmiths</a></li>
<li><a href="/wiki/Institute_of_Cancer_Research" title="Institute of Cancer Research">Institute of Cancer Research</a></li>
<li><a href="/wiki/King%27s_College_London" title="King&#39;s College London">King's College London</a></li>
<li><a href="/wiki/London_Business_School" title="London Business School">London Business School</a></li>
<li><a href="/wiki/London_School_of_Economics" title="London School of Economics">London School of Economics</a></li>
<li><a href="/wiki/London_School_of_Hygiene_%26_Tropical_Medicine" title="London School of Hygiene &amp; Tropical Medicine">London School of Hygiene &amp; Tropical Medicine</a></li>
<li><a href="/wiki/Queen_Mary_University_of_London" title="Queen Mary University of London">Queen Mary</a></li>
<li><a href="/wiki/Royal_Academy_of_Music" title="Royal Academy of Music">Royal Academy of Music</a></li>
<li><a href="/wiki/Royal_Central_School_of_Speech_and_Drama" title="Royal Central School of Speech and Drama">Royal Central School of Speech and Drama</a></li>
<li><a href="/wiki/Royal_Holloway,_University_of_London" title="Royal Holloway, University of London">Royal Holloway</a></li>
<li><a href="/wiki/Royal_Veterinary_College" title="Royal Veterinary College">Royal Veterinary College</a></li>
<li><a href="/wiki/School_of_Advanced_Study" title="School of Advanced Study">School of Advanced Study</a></li>
<li><a href="/wiki/SOAS_University_of_London" title="SOAS University of London">SOAS</a></li>
<li><a href="/wiki/University_College_London" title="University College London">University College London</a></li></ul>
</div></td></tr><tr><th scope="row" class="navbox-group" style="width:1%">Other</th><td class="navbox-list-with-group navbox-list navbox-even" style="width:100%;padding:0"><div style="padding:0 0.25em">
<ul><li><a href="/wiki/University_of_the_Arts_London" title="University of the Arts London">University of the Arts London</a></li>
<li><a href="/wiki/BPP_University" title="BPP University">BPP</a></li>
<li><a href="/wiki/University_of_East_London" title="University of East London">East London</a></li>
<li><a href="/wiki/University_of_Greenwich" title="University of Greenwich">Greenwich</a></li>
<li><a href="/wiki/Imperial_College_London" title="Imperial College London">Imperial College London</a></li>
<li><a href="/wiki/Kingston_University" title="Kingston University">Kingston</a></li>
<li><a href="/wiki/University_of_Law" title="University of Law">University of Law</a></li>
<li><a href="/wiki/London_Metropolitan_University" title="London Metropolitan University">London Met</a></li>
<li><a href="/wiki/London_South_Bank_University" title="London South Bank University">London South Bank</a></li>
<li><a href="/wiki/Middlesex_University" title="Middlesex University">Middlesex</a></li>
<li><a href="/wiki/Northeastern_University_%E2%80%93_London" title="Northeastern University – London">Northeastern University – London</a></li>
<li><a href="/wiki/Ravensbourne_University_London" title="Ravensbourne University London">Ravensbourne</a></li>
<li><a href="/wiki/Regent%27s_University_London" title="Regent&#39;s University London">Regent's University London</a></li>
<li><a href="/wiki/Richmond,_The_American_International_University_in_London" class="mw-redirect" title="Richmond, The American International University in London">Richmond</a></li>
<li><a href="/wiki/University_of_Roehampton" title="University of Roehampton">Roehampton</a></li>
<li><a href="/wiki/St_Mary%27s_University,_Twickenham" title="St Mary&#39;s University, Twickenham">St Mary's</a></li>
<li><a href="/wiki/University_of_Westminster" title="University of Westminster">Westminster</a></li>
<li><a href="/wiki/University_of_West_London" title="University of West London">West London</a></li></ul>
</div></td></tr></tbody></table><div></div></td></tr><tr><th scope="row" class="navbox-group" style="width:1%">Higher education colleges</th><td class="navbox-list-with-group navbox-list navbox-odd hlist" style="width:100%;padding:0"><div style="padding:0 0.25em">
<ul><li><a href="/wiki/Architectural_Association_School_of_Architecture" title="Architectural Association School of Architecture">Architectural Association</a></li>
<li><a href="/wiki/Bloomsbury_Institute" title="Bloomsbury Institute">Bloomsbury Institute</a></li>
<li><a href="/wiki/City_and_Guilds_of_London_Art_School" title="City and Guilds of London Art School">City and Guilds of London Art School</a></li>
<li><a href="/wiki/The_Cond%C3%A9_Nast_College_of_Fashion_%26_Design" title="The Condé Nast College of Fashion &amp; Design">The Condé Nast College of Fashion &amp; Design</a></li>
<li><a href="/wiki/Kensington_College_of_Business" title="Kensington College of Business">Kensington College of Business</a></li>
<li><a href="/wiki/London_Churchill_College" title="London Churchill College">London Churchill College</a></li>
<li><a href="/wiki/London_Interdisciplinary_School" title="London Interdisciplinary School">London Interdisciplinary School</a></li>
<li><a href="/wiki/London_School_of_Business_and_Finance" title="London School of Business and Finance">London School of Business and Finance</a></li>
<li><a href="/wiki/London_School_of_Journalism" title="London School of Journalism">London School of Journalism</a></li>
<li><a href="/wiki/Muslim_College" title="Muslim College">Muslim College</a></li>
<li><a href="/wiki/Pearson_College_London" title="Pearson College London">Pearson College</a></li>
<li><a href="/wiki/Rose_Bruford_College" title="Rose Bruford College">Rose Bruford College</a></li>
<li><a href="/wiki/Royal_College_of_Art" title="Royal College of Art">Royal College of Art</a></li>
<li><a href="/wiki/Royal_College_of_Music" title="Royal College of Music">Royal College of Music</a></li>
<li><a href="/wiki/Sotheby%27s_Institute_of_Art" title="Sotheby&#39;s Institute of Art">Sotheby's Institute of Art</a></li>
<li><a href="/wiki/South_London_Christian_College" title="South London Christian College">South London Christian College</a></li>
<li><a href="/wiki/St_Patrick%27s_College,_London" title="St Patrick&#39;s College, London">St Patrick's College, London</a></li></ul>
</div></td></tr><tr><th scope="row" class="navbox-group" style="width:1%">Further education colleges</th><td class="navbox-list-with-group navbox-list navbox-even hlist" style="width:100%;padding:0"><div style="padding:0 0.25em">
<ul><li><a href="/wiki/Ada,_the_National_College_for_Digital_Skills" title="Ada, the National College for Digital Skills">Ada</a></li>
<li><a href="/wiki/Barking_and_Dagenham_College" title="Barking and Dagenham College">Barking and Dagenham</a></li>
<li><a href="/wiki/Barnet_and_Southgate_College" title="Barnet and Southgate College">Barnet and Southgate</a></li>
<li><a href="/wiki/Capel_Manor_College" title="Capel Manor College">Capel Manor</a></li>
<li><a href="/wiki/City_and_Islington_College" title="City and Islington College">City and Islington</a></li>
<li><a href="/wiki/City_Literary_Institute" class="mw-redirect" title="City Literary Institute">City Lit</a></li>
<li><a href="/wiki/City_of_Westminster_College" title="City of Westminster College">City of Westminster</a></li>
<li><a href="/wiki/Croydon_College" title="Croydon College">Croydon</a></li>
<li><a href="/wiki/Fashion_Retail_Academy" title="Fashion Retail Academy">Fashion Retail</a></li>
<li><a href="/wiki/The_College_of_Haringey,_Enfield_and_North_East_London" title="The College of Haringey, Enfield and North East London">Haringey, Enfield and North East London</a></li>
<li><a href="/wiki/Harrow_College" title="Harrow College">Harrow</a></li>
<li><a href="/wiki/Havering_College_of_Further_and_Higher_Education" title="Havering College of Further and Higher Education">Havering</a></li>
<li><a href="/wiki/Lambeth_College" class="mw-redirect" title="Lambeth College">Lambeth</a></li>
<li><a href="/wiki/Lewisham_College" title="Lewisham College">Lewisham</a></li>
<li><a href="/wiki/London_South_East_Colleges" title="London South East Colleges">London South East Colleges</a></li>
<li><a href="/wiki/The_Marine_Society_College_of_the_Sea" title="The Marine Society College of the Sea">Marine Society</a></li>
<li><a href="/wiki/Mary_Ward_Centre" title="Mary Ward Centre">Mary Ward</a></li>
<li><a href="/wiki/Morley_College" title="Morley College">Morley</a></li>
<li><a href="/wiki/New_City_College" title="New City College">New City</a></li>
<li><a href="/wiki/Newham_College_of_Further_Education" title="Newham College of Further Education">Newham</a></li>
<li><a href="/wiki/College_of_North_West_London" title="College of North West London">North West London</a></li>
<li><a href="/wiki/Richmond_and_Hillcroft_Adult_Community_College" title="Richmond and Hillcroft Adult Community College">Richmond and Hillcroft Adult</a></li>
<li><a href="/wiki/Richmond_upon_Thames_College" title="Richmond upon Thames College">Richmond upon Thames</a></li>
<li><a href="/wiki/Southwark_College" title="Southwark College">Southwark</a></li>
<li><a href="/wiki/Sutton_College" title="Sutton College">Sutton</a></li>
<li><a href="/wiki/South_Thames_Colleges_Group" title="South Thames Colleges Group">South Thames</a></li>
<li><a href="/wiki/Stanmore_College" title="Stanmore College">Stanmore</a></li>
<li><a href="/wiki/Uxbridge_College" title="Uxbridge College">Uxbridge</a></li>
<li><a href="/wiki/Waltham_Forest_College" title="Waltham Forest College">Waltham Forest</a></li>
<li><a href="/wiki/West_London_College" title="West London College">West London</a></li>
<li><a href="/wiki/West_Thames_College" title="West Thames College">West Thames</a></li>
<li><a href="/wiki/Westminster_Kingsway_College" title="Westminster Kingsway College">Westminster Kingsway</a></li>
<li><a href="/wiki/Workers%27_Educational_Association" title="Workers&#39; Educational Association">Workers' Educational</a></li>
<li><a href="/wiki/Working_Men%27s_College" title="Working Men&#39;s College">Working Men's</a></li></ul>
</div></td></tr><tr><th scope="row" class="navbox-group" style="width:1%">Sixth form colleges</th><td class="navbox-list-with-group navbox-list navbox-odd hlist" style="width:100%;padding:0"><div style="padding:0 0.25em">
<ul><li><a href="/wiki/Big_Creative_Academy" title="Big Creative Academy">Big Creative</a></li>
<li><a href="/wiki/BSix_Sixth_Form_College" title="BSix Sixth Form College">BSix</a></li>
<li><a href="/wiki/Christ_the_King_Sixth_Form_College" title="Christ the King Sixth Form College">Christ the King</a></li>
<li><a href="/wiki/Coulsdon_Sixth_Form_College" title="Coulsdon Sixth Form College">Coulsdon</a></li>
<li><a href="/wiki/East_London_Arts_%26_Music" title="East London Arts &amp; Music">ELAM</a></li>
<li><a href="/wiki/Haringey_Sixth_Form_College" title="Haringey Sixth Form College">Haringey</a></li>
<li><a href="/wiki/Harris_Westminster_Sixth_Form" title="Harris Westminster Sixth Form">Harris Westminster</a></li>
<li><a href="/wiki/Havering_Sixth_Form_College" title="Havering Sixth Form College">Havering</a></li>
<li><a href="/wiki/John_Ruskin_College" title="John Ruskin College">John Ruskin</a></li>
<li><a href="/wiki/King%27s_College_London_Mathematics_School" title="King&#39;s College London Mathematics School">King's London Maths</a></li>
<li><a href="/wiki/Leyton_Sixth_Form_College" title="Leyton Sixth Form College">Leyton</a></li>
<li><a href="/wiki/London_Academy_of_Excellence" title="London Academy of Excellence">London Academy of Excellence</a></li>
<li><a href="/wiki/London_Academy_of_Excellence_Tottenham" title="London Academy of Excellence Tottenham">London Academy of Excellence Tottenham</a></li>
<li><a href="/wiki/Newham_Sixth_Form_College" title="Newham Sixth Form College">Newham</a></li>
<li><a href="/wiki/Newham_Collegiate_Sixth_Form_Centre" title="Newham Collegiate Sixth Form Centre">Newham Collegiate</a></li>
<li><a href="/wiki/St_Charles_Catholic_Sixth_Form_College" title="St Charles Catholic Sixth Form College">St Charles</a></li>
<li><a href="/wiki/St_Dominic%27s_Sixth_Form_College" title="St Dominic&#39;s Sixth Form College">St Dominic's</a></li>
<li><a href="/wiki/Saint_Francis_Xavier_College,_Clapham" title="Saint Francis Xavier College, Clapham">St Francis Xavier</a></li>
<li><a href="/wiki/Sir_George_Monoux_College" title="Sir George Monoux College">Sir George Monoux</a></li>
<li><a href="/wiki/South_Bank_University_Sixth_Form" title="South Bank University Sixth Form">South Bank</a></li>
<li><a href="/wiki/Tech_City_College" title="Tech City College">Tech City</a></li>
<li><a href="/wiki/William_Morris_Sixth_Form" title="William Morris Sixth Form">William Morris</a></li>
<li><a href="/wiki/Woodhouse_College" title="Woodhouse College">Woodhouse</a></li></ul>
</div></td></tr><tr><td class="navbox-abovebelow" colspan="2"><div><span class="noviewer" typeof="mw:File"><span title="List-Class article"><img alt="" src="//upload.wikimedia.org/wikipedia/en/thumb/d/db/Symbol_list_class.svg/20px-Symbol_list_class.svg.png" decoding="async" width="16" height="16" class="mw-file-element" srcset="//upload.wikimedia.org/wikipedia/en/thumb/d/db/Symbol_list_class.svg/40px-Symbol_list_class.svg.png 1.5x" data-file-width="180" data-file-height="185" /></span></span> <b><a href="/wiki/List_of_universities_and_higher_education_colleges_in_London" title="List of universities and higher education colleges in London">List</a></b></div></td></tr></tbody></table></div>
<div class="navbox-styles"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1129693374" /><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1236075235" /></div><div role="navigation" class="navbox" aria-labelledby="Universities_in_the_United_Kingdom734" style="padding:3px"><table class="nowraplinks hlist mw-collapsible mw-collapsed navbox-inner" style="border-spacing:0;background:transparent;color:inherit"><tbody><tr><th scope="col" class="navbox-title" colspan="2"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1129693374" /><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1239400231" /><div class="navbar plainlinks hlist navbar-mini"><ul><li class="nv-view"><a href="/wiki/Template:Universities_in_the_United_Kingdom" title="Template:Universities in the United Kingdom"><abbr title="View this template">v</abbr></a></li><li class="nv-talk"><a href="/wiki/Template_talk:Universities_in_the_United_Kingdom" title="Template talk:Universities in the United Kingdom"><abbr title="Discuss this template">t</abbr></a></li><li class="nv-edit"><a href="/wiki/Special:EditPage/Template:Universities_in_the_United_Kingdom" title="Special:EditPage/Template:Universities in the United Kingdom"><abbr title="Edit this template">e</abbr></a></li></ul></div><div id="Universities_in_the_United_Kingdom734" style="font-size:114%;margin:0 4em"><a href="/wiki/Universities_in_the_United_Kingdom" title="Universities in the United Kingdom">Universities in the United Kingdom</a></div></th></tr><tr><th scope="row" class="navbox-group" style="width:1%"><a href="/wiki/List_of_universities_in_England" title="List of universities in England">England</a></th><td class="navbox-list-with-group navbox-list navbox-odd" style="width:100%;padding:0"><div style="padding:0 0.25em"></div><table class="nowraplinks navbox-subgroup" style="border-spacing:0"><tbody><tr><th scope="row" class="navbox-group" style="width:1%"><a href="/wiki/East_of_England" title="East of England">East of England</a></th><td class="navbox-list-with-group navbox-list navbox-odd" style="width:100%;padding:0"><div style="padding:0 0.25em">
<ul><li><a href="/wiki/Anglia_Ruskin_University" title="Anglia Ruskin University">Anglia Ruskin</a></li>
<li><a href="/wiki/University_of_Bedfordshire" title="University of Bedfordshire">Bedfordshire</a></li>
<li><a href="/wiki/University_of_Cambridge" title="University of Cambridge">Cambridge</a></li>
<li><a href="/wiki/Cranfield_University" title="Cranfield University">Cranfield</a></li>
<li><a href="/wiki/University_of_East_Anglia" title="University of East Anglia">East Anglia</a></li>
<li><a href="/wiki/University_of_Essex" title="University of Essex">Essex</a></li>
<li><a href="/wiki/University_of_Hertfordshire" title="University of Hertfordshire">Hertfordshire</a></li>
<li><a href="/wiki/Norwich_University_of_the_Arts" title="Norwich University of the Arts">Norwich University of the Arts</a></li>
<li><a href="/wiki/University_of_Suffolk" title="University of Suffolk">Suffolk</a></li></ul>
</div></td></tr><tr><th scope="row" class="navbox-group" style="width:1%"><a href="/wiki/List_of_universities_and_higher_education_colleges_in_London" title="List of universities and higher education colleges in London">London</a></th><td class="navbox-list-with-group navbox-list navbox-odd" style="width:100%;padding:0"><div style="padding:0 0.25em"></div><table class="nowraplinks navbox-subgroup" style="border-spacing:0"><tbody><tr><th scope="row" class="navbox-group" style="width:1%"><a href="/wiki/University_of_London" title="University of London">University <br /> of London</a></th><td class="navbox-list-with-group navbox-list navbox-even" style="width:100%;padding:0"><div style="padding:0 0.25em">
<ul><li><a href="/wiki/Birkbeck,_University_of_London" title="Birkbeck, University of London">Birkbeck</a></li>
<li><a href="/wiki/Brunel_University_of_London" title="Brunel University of London">Brunel</a></li>
<li><a href="/wiki/City_St_George%27s,_University_of_London" title="City St George&#39;s, University of London">City St George's</a></li>
<li><a href="/wiki/Courtauld_Institute_of_Art" title="Courtauld Institute of Art">Courtauld</a></li>
<li><a href="/wiki/Goldsmiths,_University_of_London" title="Goldsmiths, University of London">Goldsmiths</a></li>
<li><a href="/wiki/Institute_of_Cancer_Research" title="Institute of Cancer Research">Institute of Cancer Research</a></li>
<li><a href="/wiki/King%27s_College_London" title="King&#39;s College London">King's</a></li>
<li><a href="/wiki/London_Business_School" title="London Business School">London Business School</a></li>
<li><a href="/wiki/London_School_of_Economics" title="London School of Economics">LSE</a></li>
<li><a href="/wiki/London_School_of_Hygiene_%26_Tropical_Medicine" title="London School of Hygiene &amp; Tropical Medicine">LSHTM</a></li>
<li><a href="/wiki/Queen_Mary_University_of_London" title="Queen Mary University of London">Queen Mary</a></li>
<li><a href="/wiki/Royal_Academy_of_Music" title="Royal Academy of Music">Royal Academy of Music</a></li>
<li><a href="/wiki/Royal_Central_School_of_Speech_and_Drama" title="Royal Central School of Speech and Drama">Royal Central School of Speech and Drama</a></li>
<li><a href="/wiki/Royal_Holloway,_University_of_London" title="Royal Holloway, University of London">Royal Holloway</a></li>
<li><a href="/wiki/Royal_Veterinary_College" title="Royal Veterinary College">Royal Veterinary College</a></li>
<li><a href="/wiki/School_of_Advanced_Study" title="School of Advanced Study">School of Advanced Study</a></li>
<li><a href="/wiki/SOAS_University_of_London" title="SOAS University of London">SOAS</a></li>
<li><a href="/wiki/University_College_London" title="University College London">UCL</a></li></ul>
</div></td></tr><tr><th scope="row" class="navbox-group" style="width:1%">Other</th><td class="navbox-list-with-group navbox-list navbox-odd" style="width:100%;padding:0"><div style="padding:0 0.25em">
<ul><li><a href="/wiki/University_of_East_London" title="University of East London">East London</a></li>
<li><a href="/wiki/University_of_Greenwich" title="University of Greenwich">Greenwich</a></li>
<li><a href="/wiki/Imperial_College_London" title="Imperial College London">Imperial</a></li>
<li><a href="/wiki/Kingston_University" title="Kingston University">Kingston</a></li>
<li><a href="/wiki/London_Metropolitan_University" title="London Metropolitan University">London Met</a></li>
<li><a href="/wiki/London_South_Bank_University" title="London South Bank University">London South Bank</a></li>
<li><a href="/wiki/Middlesex_University" title="Middlesex University">Middlesex</a></li>
<li><a href="/wiki/Northeastern_University_%E2%80%93_London" title="Northeastern University – London">Northeastern University – London</a></li>
<li><a href="/wiki/Ravensbourne_University_London" title="Ravensbourne University London">Ravensbourne</a></li>
<li><a href="/wiki/Regent%27s_University_London" title="Regent&#39;s University London">Regent's</a></li>
<li><a href="/wiki/Richmond,_The_American_International_University_in_London" class="mw-redirect" title="Richmond, The American International University in London">Richmond, The American International University in London</a></li>
<li><a href="/wiki/University_of_Roehampton" title="University of Roehampton">Roehampton</a></li>
<li><a href="/wiki/St_Mary%27s_University,_Twickenham" title="St Mary&#39;s University, Twickenham">St Mary's</a></li>
<li><a href="/wiki/University_of_the_Arts_London" title="University of the Arts London">University of the Arts London</a></li>
<li><a href="/wiki/University_of_Westminster" title="University of Westminster">Westminster</a></li>
<li><a href="/wiki/University_of_West_London" title="University of West London">West London</a></li></ul>
</div></td></tr></tbody></table><div></div></td></tr><tr><th scope="row" class="navbox-group" style="width:1%"><a href="/wiki/Midlands" title="Midlands">Midlands</a></th><td class="navbox-list-with-group navbox-list navbox-even" style="width:100%;padding:0"><div style="padding:0 0.25em">
<ul><li><a href="/wiki/Aston_University" title="Aston University">Aston</a></li>
<li><a href="/wiki/University_of_Birmingham" title="University of Birmingham">Birmingham</a></li>
<li><a href="/wiki/Birmingham_City_University" title="Birmingham City University">Birmingham City</a></li>
<li><a href="/wiki/Birmingham_Newman_University" title="Birmingham Newman University">Birmingham Newman</a></li>
<li><a href="/wiki/Bishop_Grosseteste_University" title="Bishop Grosseteste University">Bishop Grosseteste</a></li>
<li><a href="/wiki/Coventry_University" title="Coventry University">Coventry</a></li>
<li><a href="/wiki/De_Montfort_University" title="De Montfort University">De Montfort</a></li>
<li><a href="/wiki/University_of_Derby" title="University of Derby">Derby</a></li>
<li><a href="/wiki/Harper_Adams_University" title="Harper Adams University">Harper Adams</a></li>
<li><a href="/wiki/Keele_University" title="Keele University">Keele</a></li>
<li><a href="/wiki/University_of_Leicester" title="University of Leicester">Leicester</a></li>
<li><a href="/wiki/University_of_Lincoln" title="University of Lincoln">Lincoln</a></li>
<li><a href="/wiki/Loughborough_University" title="Loughborough University">Loughborough</a></li>
<li><a href="/wiki/University_of_Northampton" title="University of Northampton">Northampton</a></li>
<li><a href="/wiki/University_of_Nottingham" title="University of Nottingham">Nottingham</a></li>
<li><a href="/wiki/Nottingham_Trent_University" title="Nottingham Trent University">Nottingham Trent</a></li>
<li><a href="/wiki/Staffordshire_University" class="mw-redirect" title="Staffordshire University">Staffordshire</a></li>
<li><a href="/wiki/University_College_Birmingham" title="University College Birmingham">University College Birmingham</a></li>
<li><a href="/wiki/University_of_Warwick" title="University of Warwick">Warwick</a></li>
<li><a href="/wiki/University_of_Wolverhampton" title="University of Wolverhampton">Wolverhampton</a></li>
<li><a href="/wiki/University_of_Worcester" title="University of Worcester">Worcester</a></li></ul>
</div></td></tr><tr><th scope="row" class="navbox-group" style="width:1%">North</th><td class="navbox-list-with-group navbox-list navbox-odd" style="width:100%;padding:0"><div style="padding:0 0.25em">
<ul><li><a href="/wiki/University_of_Bradford" title="University of Bradford">Bradford</a></li>
<li><a href="/wiki/University_of_Central_Lancashire" class="mw-redirect" title="University of Central Lancashire">Central Lancashire</a></li>
<li><a href="/wiki/University_of_Chester" title="University of Chester">Chester</a></li>
<li><a href="/wiki/University_of_Cumbria" title="University of Cumbria">Cumbria</a></li>
<li><a href="/wiki/Durham_University" title="Durham University">Durham</a></li>
<li><a href="/wiki/Edge_Hill_University" title="Edge Hill University">Edge Hill</a></li>
<li><a href="/wiki/University_of_Greater_Manchester" title="University of Greater Manchester">Greater Manchester</a></li>
<li><a href="/wiki/University_of_Huddersfield" title="University of Huddersfield">Huddersfield</a></li>
<li><a href="/wiki/University_of_Hull" title="University of Hull">Hull</a></li>
<li><a href="/wiki/Lancaster_University" title="Lancaster University">Lancaster</a></li>
<li><a href="/wiki/University_of_Leeds" title="University of Leeds">Leeds</a></li>
<li><a href="/wiki/Leeds_Arts_University" title="Leeds Arts University">Leeds Arts</a></li>
<li><a href="/wiki/Leeds_Beckett_University" title="Leeds Beckett University">Leeds Beckett</a></li>
<li><a href="/wiki/Leeds_Trinity_University" title="Leeds Trinity University">Leeds Trinity</a></li>
<li><a href="/wiki/University_of_Liverpool" title="University of Liverpool">Liverpool</a></li>
<li><a href="/wiki/Liverpool_Hope_University" title="Liverpool Hope University">Liverpool Hope</a></li>
<li><a href="/wiki/Liverpool_John_Moores_University" title="Liverpool John Moores University">Liverpool John Moores</a></li>
<li><a href="/wiki/Liverpool_School_of_Tropical_Medicine" title="Liverpool School of Tropical Medicine">LSTM</a></li>
<li><a href="/wiki/University_of_Manchester" title="University of Manchester">Manchester</a></li>
<li><a href="/wiki/Manchester_Metropolitan_University" title="Manchester Metropolitan University">Manchester Metropolitan</a></li>
<li><a href="/wiki/Newcastle_University" title="Newcastle University">Newcastle</a></li>
<li><a href="/wiki/Northumbria_University" title="Northumbria University">Northumbria</a></li>
<li><a href="/wiki/University_of_Salford" title="University of Salford">Salford</a></li>
<li><a href="/wiki/University_of_Sheffield" title="University of Sheffield">Sheffield</a></li>
<li><a href="/wiki/Sheffield_Hallam_University" title="Sheffield Hallam University">Sheffield Hallam</a></li>
<li><a href="/wiki/University_of_Sunderland" title="University of Sunderland">Sunderland</a></li>
<li><a href="/wiki/Teesside_University" title="Teesside University">Teesside</a></li>
<li><a href="/wiki/University_of_York" title="University of York">York</a></li>
<li><a href="/wiki/York_St_John_University" title="York St John University">York St. John</a></li></ul>
</div></td></tr><tr><th scope="row" class="navbox-group" style="width:1%">South</th><td class="navbox-list-with-group navbox-list navbox-even" style="width:100%;padding:0"><div style="padding:0 0.25em">
<ul><li><a href="/wiki/Arts_University_Bournemouth" title="Arts University Bournemouth">Arts University Bournemouth</a></li>
<li><a href="/wiki/University_of_Bath" title="University of Bath">Bath</a></li>
<li><a href="/wiki/Bath_Spa_University" title="Bath Spa University">Bath Spa</a></li>
<li><a href="/wiki/Bournemouth_University" title="Bournemouth University">Bournemouth</a></li>
<li><a href="/wiki/University_of_Brighton" title="University of Brighton">Brighton</a></li>
<li><a href="/wiki/University_of_Bristol" title="University of Bristol">Bristol</a></li>
<li><a href="/wiki/University_of_Buckingham" title="University of Buckingham">Buckingham</a></li>
<li><a href="/wiki/Buckinghamshire_New_University" title="Buckinghamshire New University">Buckinghamshire New</a></li>
<li><a href="/wiki/Canterbury_Christ_Church_University" title="Canterbury Christ Church University">Canterbury Christ Church</a></li>
<li><a href="/wiki/University_of_Chichester" title="University of Chichester">Chichester</a></li>
<li><a href="/wiki/University_for_the_Creative_Arts" title="University for the Creative Arts">Creative Arts</a></li>
<li><a href="/wiki/University_of_Exeter" title="University of Exeter">Exeter</a></li>
<li><a href="/wiki/Falmouth_University" title="Falmouth University">Falmouth</a></li>
<li><a href="/wiki/Hartpury_College" class="mw-redirect" title="Hartpury College">Hartpury</a></li>
<li><a href="/wiki/University_of_Gloucestershire" title="University of Gloucestershire">Gloucestershire</a></li>
<li><a href="/wiki/University_of_Kent" title="University of Kent">Kent</a></li>
<li><a href="/wiki/University_of_Oxford" title="University of Oxford">Oxford</a></li>
<li><a href="/wiki/Oxford_Brookes_University" title="Oxford Brookes University">Oxford Brookes</a></li>
<li><a href="/wiki/University_of_Plymouth" title="University of Plymouth">Plymouth</a></li>
<li><a href="/wiki/Plymouth_Marjon_University" title="Plymouth Marjon University">Plymouth Marjon</a></li>
<li><a href="/wiki/Arts_University_Plymouth" title="Arts University Plymouth">Arts University Plymouth</a></li>
<li><a href="/wiki/University_of_Portsmouth" title="University of Portsmouth">Portsmouth</a></li>
<li><a href="/wiki/University_of_Reading" title="University of Reading">Reading</a></li>
<li><a href="/wiki/Royal_Agricultural_University" title="Royal Agricultural University">Royal Agricultural</a></li>
<li><a href="/wiki/Solent_University" class="mw-redirect" title="Solent University">Solent</a></li>
<li><a href="/wiki/University_of_Southampton" title="University of Southampton">Southampton</a></li>
<li><a href="/wiki/University_of_Surrey" title="University of Surrey">Surrey</a></li>
<li><a href="/wiki/University_of_Sussex" title="University of Sussex">Sussex</a></li>
<li><a href="/wiki/University_of_the_West_of_England" title="University of the West of England">UWE Bristol</a></li>
<li><a href="/wiki/University_of_Winchester" title="University of Winchester">Winchester</a></li></ul>
</div></td></tr></tbody></table><div></div></td></tr><tr><th scope="row" class="navbox-group" style="width:1%"><a href="/wiki/List_of_universities_in_Northern_Ireland" class="mw-redirect" title="List of universities in Northern Ireland">Northern Ireland</a></th><td class="navbox-list-with-group navbox-list navbox-odd" style="width:100%;padding:0"><div style="padding:0 0.25em">
<ul><li><a href="/wiki/Queen%27s_University_Belfast" title="Queen&#39;s University Belfast">Queen's</a></li>
<li><a href="/wiki/Ulster_University" title="Ulster University">Ulster</a></li></ul>
</div></td></tr><tr><th scope="row" class="navbox-group" style="width:1%"><a href="/wiki/List_of_universities_in_Scotland" title="List of universities in Scotland">Scotland</a></th><td class="navbox-list-with-group navbox-list navbox-even" style="width:100%;padding:0"><div style="padding:0 0.25em">
<ul><li><a href="/wiki/University_of_Aberdeen" title="University of Aberdeen">Aberdeen</a></li>
<li><a href="/wiki/Abertay_University" title="Abertay University">Abertay</a></li>
<li><a href="/wiki/University_of_Dundee" title="University of Dundee">Dundee</a></li>
<li><a href="/wiki/University_of_Edinburgh" title="University of Edinburgh">Edinburgh</a></li>
<li><a href="/wiki/Edinburgh_Napier_University" title="Edinburgh Napier University">Edinburgh Napier</a></li>
<li><a href="/wiki/University_of_Glasgow" title="University of Glasgow">Glasgow</a></li>
<li><a href="/wiki/Glasgow_Caledonian_University" title="Glasgow Caledonian University">Glasgow Caledonian</a></li>
<li><a href="/wiki/Heriot-Watt_University" title="Heriot-Watt University">Heriot-Watt</a></li>
<li><a href="/wiki/University_of_the_Highlands_and_Islands" title="University of the Highlands and Islands">Highlands and Islands</a></li>
<li><a href="/wiki/Queen_Margaret_University" title="Queen Margaret University">Queen Margaret</a></li>
<li><a href="/wiki/Robert_Gordon_University" title="Robert Gordon University">Robert Gordon</a></li>
<li><a href="/wiki/Royal_Conservatoire_of_Scotland" title="Royal Conservatoire of Scotland">Royal Conservatoire of Scotland</a></li>
<li><a href="/wiki/University_of_St_Andrews" title="University of St Andrews">St Andrews</a></li>
<li><a href="/wiki/University_of_Stirling" title="University of Stirling">Stirling</a></li>
<li><a href="/wiki/University_of_Strathclyde" title="University of Strathclyde">Strathclyde</a></li>
<li><a href="/wiki/University_of_the_West_of_Scotland" title="University of the West of Scotland">West of Scotland</a></li></ul>
</div></td></tr><tr><th scope="row" class="navbox-group" style="width:1%"><a href="/wiki/List_of_universities_in_Wales" title="List of universities in Wales">Wales</a></th><td class="navbox-list-with-group navbox-list navbox-odd" style="width:100%;padding:0"><div style="padding:0 0.25em">
<ul><li><a href="/wiki/Aberystwyth_University" title="Aberystwyth University">Aberystwyth</a></li>
<li><a href="/wiki/Bangor_University" title="Bangor University">Bangor</a></li>
<li><a href="/wiki/Cardiff_University" title="Cardiff University">Cardiff</a></li>
<li><a href="/wiki/Cardiff_Metropolitan_University" title="Cardiff Metropolitan University">Cardiff Metropolitan</a></li>
<li><a href="/wiki/University_of_South_Wales" title="University of South Wales">South Wales</a></li>
<li><a href="/wiki/Swansea_University" title="Swansea University">Swansea</a></li>
<li><a href="/wiki/University_of_Wales_Trinity_Saint_David" title="University of Wales Trinity Saint David">Wales Trinity Saint David</a></li>
<li><a href="/wiki/Wrexham_University" title="Wrexham University">Wrexham</a></li></ul>
</div></td></tr><tr><th scope="row" class="navbox-group" style="width:1%"><a href="/wiki/List_of_universities_in_the_United_Kingdom#Universities_in_British_overseas_territories" title="List of universities in the United Kingdom">Overseas territories</a></th><td class="navbox-list-with-group navbox-list navbox-even" style="width:100%;padding:0"><div style="padding:0 0.25em">
<ul><li><a href="/wiki/Bermuda_College" title="Bermuda College">Bermuda College</a></li>
<li><a href="/wiki/Cayman_Islands_Law_School" class="mw-redirect" title="Cayman Islands Law School">Cayman Islands Law School</a></li>
<li><a href="/wiki/University_of_Gibraltar" title="University of Gibraltar">Gibraltar</a></li>
<li><a href="/wiki/International_College_of_the_Cayman_Islands" title="International College of the Cayman Islands">International College of the Cayman Islands</a></li>
<li><a href="/wiki/Saint_James_School_of_Medicine" title="Saint James School of Medicine">Saint James School of Medicine</a></li>
<li><a href="/wiki/St._Matthew%27s_University" title="St. Matthew&#39;s University">St. Matthew's University</a></li>
<li><a href="/wiki/University_College_of_the_Cayman_Islands" title="University College of the Cayman Islands">University College of the Cayman Islands</a></li>
<li><a href="/wiki/University_of_Science,_Arts_and_Technology" title="University of Science, Arts and Technology">University of Science, Arts and Technology</a></li>
<li><a href="/wiki/University_of_the_West_Indies_Open_Campus" title="University of the West Indies Open Campus">University of the West Indies Open Campus</a></li></ul>
</div></td></tr><tr><th scope="row" class="navbox-group" style="width:1%"><a href="/wiki/List_of_universities_in_the_United_Kingdom#Universities_in_Crown_Dependencies" title="List of universities in the United Kingdom">Crown dependencies</a></th><td class="navbox-list-with-group navbox-list navbox-odd" style="width:100%;padding:0"><div style="padding:0 0.25em">
<ul><li><i><a href="/wiki/University_of_the_Channel_Islands_in_Guernsey" class="mw-redirect" title="University of the Channel Islands in Guernsey">University of the Channel Islands in Guernsey</a></i></li></ul>
</div></td></tr><tr><th scope="row" class="navbox-group" style="width:1%">Non-geographic</th><td class="navbox-list-with-group navbox-list navbox-even" style="width:100%;padding:0"><div style="padding:0 0.25em">
<ul><li><a href="/wiki/Arden_University" title="Arden University">Arden</a></li>
<li><a href="/wiki/BPP_University" title="BPP University">BPP</a></li>
<li><a href="/wiki/University_of_Law" title="University of Law">Law</a></li>
<li><a href="/wiki/University_of_London_International_Programmes" class="mw-redirect" title="University of London International Programmes">London International Programmes</a></li>
<li><a href="/wiki/Open_University" title="Open University">Open</a></li></ul>
</div></td></tr><tr><th scope="row" class="navbox-group" style="width:1%">Related</th><td class="navbox-list-with-group navbox-list navbox-odd" style="width:100%;padding:0"><div style="padding:0 0.25em">
<ul><li><a href="/wiki/2010_United_Kingdom_student_protests" title="2010 United Kingdom student protests">2010 United Kingdom student protests</a></li>
<li><a href="/wiki/List_of_UK_universities_by_date_of_foundation" class="mw-redirect" title="List of UK universities by date of foundation">List by date of foundation</a> <small>(<a href="/wiki/Ancient_university" title="Ancient university">Ancient</a>; <a href="/wiki/Third-oldest_university_in_England_debate" title="Third-oldest university in England debate">Third-oldest in England</a>; <a href="/wiki/Red_brick_university" class="mw-redirect" title="Red brick university">Redbrick</a>; <a href="/wiki/Plate_glass_university" title="Plate glass university">Plate glass</a>; <a href="/wiki/Post-1992_university" title="Post-1992 university">Post-1992</a>)</small></li>
<li><a href="/wiki/List_of_UK_universities_by_endowment" class="mw-redirect" title="List of UK universities by endowment">List by endowment</a></li>
<li><a href="/wiki/List_of_universities_in_the_United_Kingdom_by_enrolment" title="List of universities in the United Kingdom by enrolment">List by enrolment</a></li>
<li><a href="/wiki/Colleges_within_universities_in_the_United_Kingdom" title="Colleges within universities in the United Kingdom">Colleges within universities</a></li>
<li><a href="/wiki/British_degree_abbreviations" title="British degree abbreviations">Degree abbreviations</a></li>
<li><a href="/wiki/Higher_Education_Funding_Council_for_Wales" title="Higher Education Funding Council for Wales">HEFCW</a></li>
<li><a href="/wiki/Office_for_Students" title="Office for Students">Office for Students</a> <small>(<a href="/wiki/Higher_Education_Funding_Council_for_England" title="Higher Education Funding Council for England">HEFCE</a>)</small></li>
<li><a href="/wiki/Rankings_of_universities_in_the_United_Kingdom" title="Rankings of universities in the United Kingdom">Rankings</a></li>
<li><a href="/wiki/Scottish_Funding_Council" title="Scottish Funding Council">Scottish Funding Council</a></li>
<li><a href="/wiki/Student_loans_and_grants_in_the_United_Kingdom" title="Student loans and grants in the United Kingdom">Student loans and grants in the United Kingdom</a></li>
<li><a href="/wiki/Student_Radio_Association" title="Student Radio Association">Student Radio Association</a></li>
<li><a href="/wiki/Student_television_in_the_United_Kingdom" title="Student television in the United Kingdom">Student television in the United Kingdom</a></li>
<li><a href="/wiki/Student_unionism_in_the_United_Kingdom" title="Student unionism in the United Kingdom">Student unionism in the United Kingdom</a></li>
<li><a href="/wiki/Tuition_fees_in_the_United_Kingdom" title="Tuition fees in the United Kingdom">Tuition fees in the United Kingdom</a></li>
<li><a href="/wiki/British_undergraduate_degree_classification" title="British undergraduate degree classification">Undergraduate degree classification</a></li>
<li><a href="/wiki/UCAS" title="UCAS">UCAS</a></li>
<li><a href="/wiki/University_and_College_Union" title="University and College Union">University and College Union</a></li>
<li><a href="/wiki/Polytechnic_(United_Kingdom)" title="Polytechnic (United Kingdom)">Polytechnic</a></li></ul>
</div></td></tr><tr><td class="navbox-abovebelow" colspan="2"><div>
<ul><li><span class="noviewer" typeof="mw:File"><span title="Category"><img alt="" src="//upload.wikimedia.org/wikipedia/en/thumb/9/96/Symbol_category_class.svg/20px-Symbol_category_class.svg.png" decoding="async" width="16" height="16" class="mw-file-element" srcset="//upload.wikimedia.org/wikipedia/en/thumb/9/96/Symbol_category_class.svg/40px-Symbol_category_class.svg.png 1.5x" data-file-width="180" data-file-height="185" /></span></span> <a href="/wiki/Category:Universities_in_the_United_Kingdom" title="Category:Universities in the United Kingdom">Category</a></li>
<li><span class="noviewer" typeof="mw:File"><span title="List-Class article"><img alt="" src="//upload.wikimedia.org/wikipedia/en/thumb/d/db/Symbol_list_class.svg/20px-Symbol_list_class.svg.png" decoding="async" width="16" height="16" class="mw-file-element" srcset="//upload.wikimedia.org/wikipedia/en/thumb/d/db/Symbol_list_class.svg/40px-Symbol_list_class.svg.png 1.5x" data-file-width="180" data-file-height="185" /></span></span> <a href="/wiki/List_of_universities_in_the_United_Kingdom" title="List of universities in the United Kingdom">List</a></li></ul>
</div></td></tr></tbody></table></div>
<div class="navbox-styles"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1129693374" /><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1236075235" /></div><div role="navigation" class="navbox" aria-labelledby="30px_Venues_of_the_1908_Summer_Olympics_(London)135" style="padding:3px"><table class="nowraplinks mw-collapsible autocollapse navbox-inner" style="border-spacing:0;background:transparent;color:inherit"><tbody><tr><th scope="col" class="navbox-title" colspan="2"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1129693374" /><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1239400231" /><div class="navbar plainlinks hlist navbar-mini"><ul><li class="nv-view"><a href="/wiki/Template:1908_Summer_Olympics_venues" title="Template:1908 Summer Olympics venues"><abbr title="View this template">v</abbr></a></li><li class="nv-talk"><a href="/wiki/Template_talk:1908_Summer_Olympics_venues" title="Template talk:1908 Summer Olympics venues"><abbr title="Discuss this template">t</abbr></a></li><li class="nv-edit"><a href="/wiki/Special:EditPage/Template:1908_Summer_Olympics_venues" title="Special:EditPage/Template:1908 Summer Olympics venues"><abbr title="Edit this template">e</abbr></a></li></ul></div><div id="30px_Venues_of_the_1908_Summer_Olympics_(London)135" style="font-size:114%;margin:0 4em"><span typeof="mw:File"><a href="/wiki/File:Olympic_rings_without_rims.svg" class="mw-file-description"><img src="//upload.wikimedia.org/wikipedia/commons/thumb/5/5c/Olympic_rings_without_rims.svg/40px-Olympic_rings_without_rims.svg.png" decoding="async" width="30" height="14" class="mw-file-element" srcset="//upload.wikimedia.org/wikipedia/commons/thumb/5/5c/Olympic_rings_without_rims.svg/60px-Olympic_rings_without_rims.svg.png 1.5x" data-file-width="342" data-file-height="158" /></a></span> <a href="/wiki/Venues_of_the_1908_Summer_Olympics" class="mw-redirect" title="Venues of the 1908 Summer Olympics">Venues</a> of the <a href="/wiki/1908_Summer_Olympics" title="1908 Summer Olympics">1908 Summer Olympics</a> (<a href="/wiki/London" title="London">London</a>)</div></th></tr><tr><td colspan="2" class="navbox-list navbox-odd hlist" style="width:100%;padding:0"><div style="padding:0 0.25em">
<ul><li><a href="/wiki/All_England_Lawn_Tennis_and_Croquet_Club" title="All England Lawn Tennis and Croquet Club">All England Lawn Tennis and Croquet Club</a></li>
<li><a href="/wiki/Bisley,_Surrey" title="Bisley, Surrey">Bisley Ranges</a></li>
<li><a href="/wiki/Franco-British_Exhibition" title="Franco-British Exhibition">Franco-British Exhibition Fencing Grounds</a></li>
<li><a href="/wiki/Henley_Royal_Regatta" title="Henley Royal Regatta">Henley Royal Regatta</a></li>
<li><a href="/wiki/Hurlingham_Club" class="mw-redirect" title="Hurlingham Club">Hurlingham Club</a></li>
<li><a class="mw-selflink selflink">Northampton Institute</a></li>
<li><a href="/wiki/Prince%27s_Skating_Club" title="Prince&#39;s Skating Club">Prince's Skating Club</a></li>
<li><a href="/wiki/Queen%27s_Club" title="Queen&#39;s Club">Queen's Club</a></li>
<li><a href="/wiki/The_Solent" title="The Solent">Solent</a></li>
<li><a href="/wiki/Southampton_Water" title="Southampton Water">Southampton Water</a></li>
<li><a href="/wiki/Uxendon_Shooting_School_Club" title="Uxendon Shooting School Club">Uxendon Shooting School Club</a></li>
<li><a href="/wiki/White_City_Stadium" title="White City Stadium">White City Stadium</a></li></ul>
</div></td></tr></tbody></table></div>
<div class="navbox-styles"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1129693374" /><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1236075235" /></div><div role="navigation" class="navbox" aria-labelledby="Olympic_venues_in_boxing95" style="padding:3px"><table class="nowraplinks mw-collapsible autocollapse navbox-inner" style="border-spacing:0;background:transparent;color:inherit"><tbody><tr><th scope="col" class="navbox-title" colspan="3"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1129693374" /><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1239400231" /><div class="navbar plainlinks hlist navbar-mini"><ul><li class="nv-view"><a href="/wiki/Template:Olympic_venues_boxing" title="Template:Olympic venues boxing"><abbr title="View this template">v</abbr></a></li><li class="nv-talk"><a href="/wiki/Template_talk:Olympic_venues_boxing" title="Template talk:Olympic venues boxing"><abbr title="Discuss this template">t</abbr></a></li><li class="nv-edit"><a href="/wiki/Special:EditPage/Template:Olympic_venues_boxing" title="Special:EditPage/Template:Olympic venues boxing"><abbr title="Edit this template">e</abbr></a></li></ul></div><div id="Olympic_venues_in_boxing95" style="font-size:114%;margin:0 4em"><a href="/wiki/List_of_Olympic_venues_in_boxing" title="List of Olympic venues in boxing">Olympic venues in boxing</a></div></th></tr><tr><th scope="row" class="navbox-group" style="width:1%">20th&#160;century</th><td class="navbox-list-with-group navbox-list navbox-odd hlist" style="width:100%;padding:0"><div style="padding:0 0.25em">
<ul><li><a href="/wiki/Boxing_at_the_1904_Summer_Olympics" title="Boxing at the 1904 Summer Olympics">1904</a>: <a href="/wiki/Francis_Gymnasium_(St._Louis)" class="mw-redirect" title="Francis Gymnasium (St. Louis)">Francis Gymnasium</a></li>
<li><a href="/wiki/Boxing_at_the_1908_Summer_Olympics" title="Boxing at the 1908 Summer Olympics">1908</a>: <a class="mw-selflink selflink">Northampton Institute</a></li>
<li><a href="/wiki/Boxing_at_the_1920_Summer_Olympics" title="Boxing at the 1920 Summer Olympics">1920</a>: <a href="/wiki/Antwerp_Zoo" title="Antwerp Zoo">Antwerp Zoo</a></li>
<li><a href="/wiki/Boxing_at_the_1924_Summer_Olympics" title="Boxing at the 1924 Summer Olympics">1924</a>: <a href="/wiki/V%C3%A9lodrome_d%27hiver" class="mw-redirect" title="Vélodrome d&#39;hiver">Vélodrome d'hiver</a></li>
<li><a href="/wiki/Boxing_at_the_1928_Summer_Olympics" title="Boxing at the 1928 Summer Olympics">1928</a>: <a href="/wiki/Krachtsportgebouw" title="Krachtsportgebouw">Krachtsportgebouw</a></li>
<li><a href="/wiki/Boxing_at_the_1932_Summer_Olympics" title="Boxing at the 1932 Summer Olympics">1932</a>: <a href="/wiki/Grand_Olympic_Auditorium" title="Grand Olympic Auditorium">Olympic Auditorium</a></li>
<li><a href="/wiki/Boxing_at_the_1936_Summer_Olympics" title="Boxing at the 1936 Summer Olympics">1936</a>: <a href="/wiki/Deutschlandhalle" title="Deutschlandhalle">Deutschlandhalle</a></li>
<li><a href="/wiki/Boxing_at_the_1948_Summer_Olympics" title="Boxing at the 1948 Summer Olympics">1948</a>: <a href="/wiki/Wembley_Arena" title="Wembley Arena">Empire Pool</a>, <a href="/wiki/Earls_Court_Exhibition_Centre" title="Earls Court Exhibition Centre">Empress Hall, Earl's Court</a></li>
<li><a href="/wiki/Boxing_at_the_1952_Summer_Olympics" title="Boxing at the 1952 Summer Olympics">1952</a>: <a href="/wiki/T%C3%B6%C3%B6l%C3%B6_Sports_Hall" title="Töölö Sports Hall">Messuhalli</a></li>
<li><a href="/wiki/Boxing_at_the_1956_Summer_Olympics" title="Boxing at the 1956 Summer Olympics">1956</a>: <a href="/wiki/Festival_Hall,_Melbourne" class="mw-redirect" title="Festival Hall, Melbourne">West Melbourne Stadium</a></li>
<li><a href="/wiki/Boxing_at_the_1960_Summer_Olympics" title="Boxing at the 1960 Summer Olympics">1960</a>: <a href="/wiki/PalaLottomatica" title="PalaLottomatica">Palazzo dello Sport</a></li>
<li><a href="/wiki/Boxing_at_the_1964_Summer_Olympics" title="Boxing at the 1964 Summer Olympics">1964</a>: <a href="/wiki/Korakuen_Hall" title="Korakuen Hall">Korakuen Ice Palace</a></li>
<li><a href="/wiki/Boxing_at_the_1968_Summer_Olympics" title="Boxing at the 1968 Summer Olympics">1968</a>: <a href="/wiki/Arena_M%C3%A9xico" title="Arena México">Arena México</a></li>
<li><a href="/wiki/Boxing_at_the_1972_Summer_Olympics" title="Boxing at the 1972 Summer Olympics">1972</a>: <a href="/wiki/Boxhalle" title="Boxhalle">Boxhalle</a></li>
<li><a href="/wiki/Boxing_at_the_1976_Summer_Olympics" title="Boxing at the 1976 Summer Olympics">1976</a>: <a href="/wiki/Maurice_Richard_Arena" title="Maurice Richard Arena">Maurice Richard Arena</a>, <a href="/wiki/Montreal_Forum" title="Montreal Forum">Montreal Forum</a> (finals)</li>
<li><a href="/wiki/Boxing_at_the_1980_Summer_Olympics" title="Boxing at the 1980 Summer Olympics">1980</a>: <a href="/wiki/Olympic_Stadium_(Moscow)" title="Olympic Stadium (Moscow)">Indoor Stadium</a></li>
<li><a href="/wiki/Boxing_at_the_1984_Summer_Olympics" title="Boxing at the 1984 Summer Olympics">1984</a>: <a href="/wiki/Los_Angeles_Memorial_Sports_Arena" title="Los Angeles Memorial Sports Arena">Los Angeles Memorial Sports Arena</a></li>
<li><a href="/wiki/Boxing_at_the_1988_Summer_Olympics" title="Boxing at the 1988 Summer Olympics">1988</a>: <a href="/wiki/Jamsil_Students%27_Gymnasium" title="Jamsil Students&#39; Gymnasium">Jamsil Students' Gymnasium</a></li>
<li><a href="/wiki/Boxing_at_the_1992_Summer_Olympics" title="Boxing at the 1992 Summer Olympics">1992</a>: <a href="/wiki/Pavell%C3%B3_Club_Joventut_Badalona" title="Pavelló Club Joventut Badalona">Pavelló Club Joventut Badalona</a></li>
<li><a href="/wiki/Boxing_at_the_1996_Summer_Olympics" title="Boxing at the 1996 Summer Olympics">1996</a>: <a href="/wiki/McCamish_Pavilion" title="McCamish Pavilion">Alexander Memorial Coliseum</a></li></ul>
</div></td><td class="noviewer navbox-image" rowspan="2" style="width:1px;padding:0 0 0 2px"><div><span typeof="mw:File"><a href="/wiki/File:Boxing_pictogram.svg" class="mw-file-description"><img src="//upload.wikimedia.org/wikipedia/commons/thumb/c/c2/Boxing_pictogram.svg/60px-Boxing_pictogram.svg.png" decoding="async" width="50" height="50" class="mw-file-element" srcset="//upload.wikimedia.org/wikipedia/commons/thumb/c/c2/Boxing_pictogram.svg/120px-Boxing_pictogram.svg.png 1.5x" data-file-width="300" data-file-height="300" /></a></span></div></td></tr><tr><th scope="row" class="navbox-group" style="width:1%">21st&#160;century</th><td class="navbox-list-with-group navbox-list navbox-even hlist" style="width:100%;padding:0"><div style="padding:0 0.25em">
<ul><li><a href="/wiki/Boxing_at_the_2000_Summer_Olympics" title="Boxing at the 2000 Summer Olympics">2000</a>: <a href="/wiki/Sydney_Convention_%26_Exhibition_Centre" class="mw-redirect" title="Sydney Convention &amp; Exhibition Centre">Sydney Convention &amp; Exhibition Centre</a></li>
<li><a href="/wiki/Boxing_at_the_2004_Summer_Olympics" title="Boxing at the 2004 Summer Olympics">2004</a>: <a href="/wiki/Peristeri_Olympic_Boxing_Hall" title="Peristeri Olympic Boxing Hall">Peristeri Olympic Boxing Hall</a></li>
<li><a href="/wiki/Boxing_at_the_2008_Summer_Olympics" title="Boxing at the 2008 Summer Olympics">2008</a>: <a href="/wiki/Workers_Indoor_Arena" class="mw-redirect" title="Workers Indoor Arena">Workers Indoor Arena</a></li>
<li><a href="/wiki/Boxing_at_the_2012_Summer_Olympics" title="Boxing at the 2012 Summer Olympics">2012</a>: <a href="/wiki/ExCeL_London" class="mw-redirect" title="ExCeL London">ExCeL</a></li>
<li><a href="/wiki/Boxing_at_the_2016_Summer_Olympics" title="Boxing at the 2016 Summer Olympics">2016</a>: <a href="/wiki/Riocentro" title="Riocentro">Riocentro</a> – Pavilion 6</li>
<li><a href="/wiki/Boxing_at_the_2020_Summer_Olympics" title="Boxing at the 2020 Summer Olympics">2020</a>: <a href="/wiki/Ry%C5%8Dgoku_Kokugikan" title="Ryōgoku Kokugikan">Kokugikan Arena</a></li>
<li><a href="/wiki/Boxing_at_the_2024_Summer_Olympics" title="Boxing at the 2024 Summer Olympics">2024</a>: <a href="/wiki/Parc_des_Expositions_de_Villepinte" title="Parc des Expositions de Villepinte">Arena Paris Nord</a>, <a href="/wiki/Stade_Roland_Garros" title="Stade Roland Garros">Stade Roland Garros</a> (finals)</li>
<li><a href="/wiki/2028_Summer_Olympics" title="2028 Summer Olympics">2028</a>: <a href="/wiki/Peacock_Theater" title="Peacock Theater">LA Live Theater</a>, <a href="/wiki/Crypto.com_Arena" title="Crypto.com Arena">Crypto.com Arena</a></li>
<li><a href="/wiki/2032_Summer_Olympics" title="2032 Summer Olympics">2032</a>: <a href="/wiki/Moreton_Bay_Indoor_Sports_Centre" class="mw-redirect" title="Moreton Bay Indoor Sports Centre">Moreton Bay Indoor Sports Centre</a></li></ul>
</div></td></tr></tbody></table></div></div></td></tr></tbody></table></div>
<div class="navbox-styles"><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1129693374" /><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1236075235" /><style data-mw-deduplicate="TemplateStyles:r1038841319">.mw-parser-output .tooltip-dotted{border-bottom:1px dotted;cursor:help}</style></div><div role="navigation" class="navbox authority-control" aria-labelledby="Authority_control_databases_frameless&amp;#124;text-top&amp;#124;10px&amp;#124;alt=Edit_this_at_Wikidata&amp;#124;link=https&amp;#58;//www.wikidata.org/wiki/Q1094046#identifiers&amp;#124;class=noprint&amp;#124;Edit_this_at_Wikidata1486" style="padding:3px"><table class="nowraplinks hlist mw-collapsible autocollapse navbox-inner" style="border-spacing:0;background:transparent;color:inherit"><tbody><tr><th scope="col" class="navbox-title" colspan="2"><div id="Authority_control_databases_frameless&amp;#124;text-top&amp;#124;10px&amp;#124;alt=Edit_this_at_Wikidata&amp;#124;link=https&amp;#58;//www.wikidata.org/wiki/Q1094046#identifiers&amp;#124;class=noprint&amp;#124;Edit_this_at_Wikidata1486" style="font-size:114%;margin:0 4em"><a href="/wiki/Help:Authority_control" title="Help:Authority control">Authority control databases</a> <span class="mw-valign-text-top noprint" typeof="mw:File/Frameless"><a href="https://www.wikidata.org/wiki/Q1094046#identifiers" title="Edit this at Wikidata"><img alt="Edit this at Wikidata" src="//upload.wikimedia.org/wikipedia/en/thumb/8/8a/OOjs_UI_icon_edit-ltr-progressive.svg/20px-OOjs_UI_icon_edit-ltr-progressive.svg.png" decoding="async" width="10" height="10" class="mw-file-element" data-file-width="20" data-file-height="20" /></a></span></div></th></tr><tr><th scope="row" class="navbox-group" style="width:1%">International</th><td class="navbox-list-with-group navbox-list navbox-odd" style="width:100%;padding:0"><div style="padding:0 0.25em"><ul><li><span class="uid"><a rel="nofollow" class="external text" href="https://isni.org/isni/0000000419368497">ISNI</a></span></li><li><span class="uid"><a rel="nofollow" class="external text" href="https://viaf.org/viaf/169673265">VIAF</a></span></li></ul></div></td></tr><tr><th scope="row" class="navbox-group" style="width:1%">National</th><td class="navbox-list-with-group navbox-list navbox-even" style="width:100%;padding:0"><div style="padding:0 0.25em"><ul><li><span class="uid"><a rel="nofollow" class="external text" href="https://d-nb.info/gnd/276111-7">Germany</a></span></li><li><span class="uid"><a rel="nofollow" class="external text" href="https://id.loc.gov/authorities/n50041710">United States</a></span></li><li><span class="uid"><span class="rt-commentedText tooltip tooltip-dotted" title="City University (Londýn, Anglie)"><a rel="nofollow" class="external text" href="https://aleph.nkp.cz/F/?func=find-c&amp;local_base=aut&amp;ccl_term=ica=ko2015865187&amp;CON_LNG=ENG">Czech Republic</a></span></span></li><li><span class="uid"><a rel="nofollow" class="external text" href="https://www.nli.org.il/en/authorities/987007605581805171">Israel</a></span></li></ul></div></td></tr><tr><th scope="row" class="navbox-group" style="width:1%">Geographic</th><td class="navbox-list-with-group navbox-list navbox-odd" style="width:100%;padding:0"><div style="padding:0 0.25em"><ul><li><span class="uid"><a rel="nofollow" class="external text" href="https://musicbrainz.org/place/25a29860-1690-43b7-a54a-88881b8f559c">MusicBrainz place</a></span></li></ul></div></td></tr><tr><th scope="row" class="navbox-group" style="width:1%">People</th><td class="navbox-list-with-group navbox-list navbox-even" style="width:100%;padding:0"><div style="padding:0 0.25em"><ul><li><span class="uid"><a href="/wiki/International_Standard_Identifier_for_Libraries_and_Related_Organizations" title="International Standard Identifier for Libraries and Related Organizations">ISIL</a>: GB-UkLoCU</span></li><li><span class="uid"><a rel="nofollow" class="external text" href="https://trove.nla.gov.au/people/1237664">Trove</a></span></li></ul></div></td></tr><tr><th scope="row" class="navbox-group" style="width:1%">Other</th><td class="navbox-list-with-group navbox-list navbox-odd" style="width:100%;padding:0"><div style="padding:0 0.25em"><ul><li><span class="uid"><a rel="nofollow" class="external text" href="https://www.idref.fr/027941582">IdRef</a></span></li></ul></div></td></tr></tbody></table></div>
<!-- 
NewPP limit report
Parsed by mw‐web.codfw.canary‐86b8f7b789‐tcc2m
Cached time: 20250523235530
Cache expiry: 2592000
Reduced expiry: false
Complications: [vary‐revision‐sha1, show‐toc]
CPU time usage: 1.237 seconds
Real time usage: 1.487 seconds
Preprocessor visited node count: 6554/1000000
Revision size: 52124/2097152 bytes
Post‐expand include size: 389179/2097152 bytes
Template argument size: 84459/2097152 bytes
Highest expansion depth: 17/100
Expensive parser function count: 8/500
Unstrip recursion depth: 1/20
Unstrip post‐expand size: 289041/5000000 bytes
Lua time usage: 0.740/10.000 seconds
Lua memory usage: 8078683/52428800 bytes
Number of Wikibase entities loaded: 1/500
-->
<!--
Transclusion expansion time report (%,ms,calls,template)
100.00% 1185.956      1 -total
 44.95%  533.058      1 Template:Reflist
 35.50%  421.033     54 Template:Cite_web
 13.71%  162.617     11 Template:Navbox
  9.99%  118.445      1 Template:Infobox_university
  9.77%  115.869      2 Template:Infobox
  8.74%  103.647      1 Template:City,_University_of_London
  5.98%   70.896      1 Template:Navboxes
  5.67%   67.241      1 Template:Short_description
  4.77%   56.576      9 Template:Cite_news
-->

<!-- Saved in parser cache with key enwiki:pcache:533484:|#|:idhash:canonical and timestamp 20250523235530 and revision id 1283711199. Rendering was triggered because: page-view
 -->
</div><!--esi <esi:include src="/esitest-fa8a495983347898/content" /> --><noscript><img src="https://en.wikipedia.org/wiki/Special:CentralAutoLogin/start?type=1x1&amp;usesul3=1" alt="" width="1" height="1" style="border: none; position: absolute;"></noscript>
<div class="printfooter" data-nosnippet="">Retrieved from "<a dir="ltr" href="https://en.wikipedia.org/w/index.php?title=City,_University_of_London&amp;oldid=1283711199">https://en.wikipedia.org/w/index.php?title=City,_University_of_London&amp;oldid=1283711199</a>"</div></div>
					<div id="catlinks" class="catlinks" data-mw="interface"><div id="mw-normal-catlinks" class="mw-normal-catlinks"><a href="/wiki/Help:Category" title="Help:Category">Categories</a>: <ul><li><a href="/wiki/Category:City,_University_of_London" title="Category:City, University of London">City, University of London</a></li><li><a href="/wiki/Category:Optometry_schools" title="Category:Optometry schools">Optometry schools</a></li><li><a href="/wiki/Category:Schools_of_informatics" title="Category:Schools of informatics">Schools of informatics</a></li><li><a href="/wiki/Category:Universities_and_colleges_established_in_1894" title="Category:Universities and colleges established in 1894">Universities and colleges established in 1894</a></li><li><a href="/wiki/Category:1894_establishments_in_England" title="Category:1894 establishments in England">1894 establishments in England</a></li><li><a href="/wiki/Category:Venues_of_the_1908_Summer_Olympics" title="Category:Venues of the 1908 Summer Olympics">Venues of the 1908 Summer Olympics</a></li><li><a href="/wiki/Category:Olympic_boxing_venues" title="Category:Olympic boxing venues">Olympic boxing venues</a></li><li><a href="/wiki/Category:Universities_UK" title="Category:Universities UK">Universities UK</a></li></ul></div><div id="mw-hidden-catlinks" class="mw-hidden-catlinks mw-hidden-cats-hidden">Hidden categories: <ul><li><a href="/wiki/Category:Pages_using_gadget_WikiMiniAtlas" title="Category:Pages using gadget WikiMiniAtlas">Pages using gadget WikiMiniAtlas</a></li><li><a href="/wiki/Category:All_articles_with_dead_external_links" title="Category:All articles with dead external links">All articles with dead external links</a></li><li><a href="/wiki/Category:Articles_with_dead_external_links_from_November_2017" title="Category:Articles with dead external links from November 2017">Articles with dead external links from November 2017</a></li><li><a href="/wiki/Category:Articles_with_permanently_dead_external_links" title="Category:Articles with permanently dead external links">Articles with permanently dead external links</a></li><li><a href="/wiki/Category:Articles_with_short_description" title="Category:Articles with short description">Articles with short description</a></li><li><a href="/wiki/Category:Short_description_is_different_from_Wikidata" title="Category:Short description is different from Wikidata">Short description is different from Wikidata</a></li><li><a href="/wiki/Category:Use_dmy_dates_from_December_2019" title="Category:Use dmy dates from December 2019">Use dmy dates from December 2019</a></li><li><a href="/wiki/Category:Articles_to_be_merged_from_March_2025" title="Category:Articles to be merged from March 2025">Articles to be merged from March 2025</a></li><li><a href="/wiki/Category:All_articles_to_be_merged" title="Category:All articles to be merged">All articles to be merged</a></li><li><a href="/wiki/Category:Coordinates_on_Wikidata" title="Category:Coordinates on Wikidata">Coordinates on Wikidata</a></li><li><a href="/wiki/Category:Articles_using_infobox_university" title="Category:Articles using infobox university">Articles using infobox university</a></li><li><a href="/wiki/Category:Pages_using_infobox_university_with_the_affiliations_parameter" title="Category:Pages using infobox university with the affiliations parameter">Pages using infobox university with the affiliations parameter</a></li><li><a href="/wiki/Category:All_articles_with_unsourced_statements" title="Category:All articles with unsourced statements">All articles with unsourced statements</a></li><li><a href="/wiki/Category:Articles_with_unsourced_statements_from_October_2023" title="Category:Articles with unsourced statements from October 2023">Articles with unsourced statements from October 2023</a></li><li><a href="/wiki/Category:Articles_with_unsourced_statements_from_December_2010" title="Category:Articles with unsourced statements from December 2010">Articles with unsourced statements from December 2010</a></li><li><a href="/wiki/Category:Commons_category_link_from_Wikidata" title="Category:Commons category link from Wikidata">Commons category link from Wikidata</a></li><li><a href="/wiki/Category:Webarchive_template_wayback_links" title="Category:Webarchive template wayback links">Webarchive template wayback links</a></li></ul></div></div>
				</div>
			</main>
			
		</div>
		<div class="mw-footer-container">
			
<footer id="footer" class="mw-footer" >
	<ul id="footer-info">
	<li id="footer-info-lastmod"> This page was last edited on 3 April 2025, at 05:31<span class="anonymous-show">&#160;(UTC)</span>.</li>
	<li id="footer-info-copyright">Text is available under the <a href="/wiki/Wikipedia:Text_of_the_Creative_Commons_Attribution-ShareAlike_4.0_International_License" title="Wikipedia:Text of the Creative Commons Attribution-ShareAlike 4.0 International License">Creative Commons Attribution-ShareAlike 4.0 License</a>;
additional terms may apply. By using this site, you agree to the <a href="https://foundation.wikimedia.org/wiki/Special:MyLanguage/Policy:Terms_of_Use" class="extiw" title="foundation:Special:MyLanguage/Policy:Terms of Use">Terms of Use</a> and <a href="https://foundation.wikimedia.org/wiki/Special:MyLanguage/Policy:Privacy_policy" class="extiw" title="foundation:Special:MyLanguage/Policy:Privacy policy">Privacy Policy</a>. Wikipedia® is a registered trademark of the <a rel="nofollow" class="external text" href="https://wikimediafoundation.org/">Wikimedia Foundation, Inc.</a>, a non-profit organization.</li>
</ul>

	<ul id="footer-places">
	<li id="footer-places-privacy"><a href="https://foundation.wikimedia.org/wiki/Special:MyLanguage/Policy:Privacy_policy">Privacy policy</a></li>
	<li id="footer-places-about"><a href="/wiki/Wikipedia:About">About Wikipedia</a></li>
	<li id="footer-places-disclaimers"><a href="/wiki/Wikipedia:General_disclaimer">Disclaimers</a></li>
	<li id="footer-places-contact"><a href="//en.wikipedia.org/wiki/Wikipedia:Contact_us">Contact Wikipedia</a></li>
	<li id="footer-places-wm-codeofconduct"><a href="https://foundation.wikimedia.org/wiki/Special:MyLanguage/Policy:Universal_Code_of_Conduct">Code of Conduct</a></li>
	<li id="footer-places-developers"><a href="https://developer.wikimedia.org">Developers</a></li>
	<li id="footer-places-statslink"><a href="https://stats.wikimedia.org/#/en.wikipedia.org">Statistics</a></li>
	<li id="footer-places-cookiestatement"><a href="https://foundation.wikimedia.org/wiki/Special:MyLanguage/Policy:Cookie_statement">Cookie statement</a></li>
	<li id="footer-places-mobileview"><a href="//en.m.wikipedia.org/w/index.php?title=City,_University_of_London&amp;mobileaction=toggle_view_mobile" class="noprint stopMobileRedirectToggle">Mobile view</a></li>
</ul>

	<ul id="footer-icons" class="noprint">
	<li id="footer-copyrightico"><a href="https://www.wikimedia.org/" class="cdx-button cdx-button--fake-button cdx-button--size-large cdx-button--fake-button--enabled"><picture><source media="(min-width: 500px)" srcset="/static/images/footer/wikimedia-button.svg" width="84" height="29"><img src="/static/images/footer/wikimedia.svg" width="25" height="25" alt="Wikimedia Foundation" lang="en" loading="lazy"></picture></a></li>
	<li id="footer-poweredbyico"><a href="https://www.mediawiki.org/" class="cdx-button cdx-button--fake-button cdx-button--size-large cdx-button--fake-button--enabled"><picture><source media="(min-width: 500px)" srcset="/w/resources/assets/poweredby_mediawiki.svg" width="88" height="31"><img src="/w/resources/assets/mediawiki_compact.svg" alt="Powered by MediaWiki" lang="en" width="25" height="25" loading="lazy"></picture></a></li>
</ul>

</footer>

		</div>
	</div> 
</div> 
<div class="vector-header-container vector-sticky-header-container no-font-mode-scale">
	<div id="vector-sticky-header" class="vector-sticky-header">
		<div class="vector-sticky-header-start">
			<div class="vector-sticky-header-icon-start vector-button-flush-left vector-button-flush-right" aria-hidden="true">
				<button class="cdx-button cdx-button--weight-quiet cdx-button--icon-only vector-sticky-header-search-toggle" tabindex="-1" data-event-name="ui.vector-sticky-search-form.icon"><span class="vector-icon mw-ui-icon-search mw-ui-icon-wikimedia-search"></span>

<span>Search</span>
			</button>
		</div>
			
		<div role="search" class="vector-search-box-vue  vector-search-box-show-thumbnail vector-search-box">
			<div class="vector-typeahead-search-container">
				<div class="cdx-typeahead-search cdx-typeahead-search--show-thumbnail">
					<form action="/w/index.php" id="vector-sticky-search-form" class="cdx-search-input cdx-search-input--has-end-button">
						<div  class="cdx-search-input__input-wrapper"  data-search-loc="header-moved">
							<div class="cdx-text-input cdx-text-input--has-start-icon">
								<input
									class="cdx-text-input__input mw-searchInput"
									
									type="search" name="search" placeholder="Search Wikipedia">
								<span class="cdx-text-input__icon cdx-text-input__start-icon"></span>
							</div>
							<input type="hidden" name="title" value="Special:Search">
						</div>
						<button class="cdx-button cdx-search-input__end-button">Search</button>
					</form>
				</div>
			</div>
		</div>
		<div class="vector-sticky-header-context-bar">
				<nav aria-label="Contents" class="vector-toc-landmark">
						
					<div id="vector-sticky-header-toc" class="vector-dropdown mw-portlet mw-portlet-sticky-header-toc vector-sticky-header-toc vector-button-flush-left"  >
						<input type="checkbox" id="vector-sticky-header-toc-checkbox" role="button" aria-haspopup="true" data-event-name="ui.dropdown-vector-sticky-header-toc" class="vector-dropdown-checkbox "  aria-label="Toggle the table of contents"  >
						<label id="vector-sticky-header-toc-label" for="vector-sticky-header-toc-checkbox" class="vector-dropdown-label cdx-button cdx-button--fake-button cdx-button--fake-button--enabled cdx-button--weight-quiet cdx-button--icon-only " aria-hidden="true"  ><span class="vector-icon mw-ui-icon-listBullet mw-ui-icon-wikimedia-listBullet"></span>

<span class="vector-dropdown-label-text">Toggle the table of contents</span>
						</label>
						<div class="vector-dropdown-content">
					
						<div id="vector-sticky-header-toc-unpinned-container" class="vector-unpinned-container">
						</div>
					
						</div>
					</div>
			</nav>
				<div class="vector-sticky-header-context-bar-primary" aria-hidden="true" ><span class="mw-page-title-main">City, University of London</span></div>
			</div>
		</div>
		<div class="vector-sticky-header-end" aria-hidden="true">
			<div class="vector-sticky-header-icons">
				<a href="#" class="cdx-button cdx-button--fake-button cdx-button--fake-button--enabled cdx-button--weight-quiet cdx-button--icon-only" id="ca-talk-sticky-header" tabindex="-1" data-event-name="talk-sticky-header"><span class="vector-icon mw-ui-icon-speechBubbles mw-ui-icon-wikimedia-speechBubbles"></span>

<span></span>
			</a>
			<a href="#" class="cdx-button cdx-button--fake-button cdx-button--fake-button--enabled cdx-button--weight-quiet cdx-button--icon-only" id="ca-subject-sticky-header" tabindex="-1" data-event-name="subject-sticky-header"><span class="vector-icon mw-ui-icon-article mw-ui-icon-wikimedia-article"></span>

<span></span>
			</a>
			<a href="#" class="cdx-button cdx-button--fake-button cdx-button--fake-button--enabled cdx-button--weight-quiet cdx-button--icon-only" id="ca-history-sticky-header" tabindex="-1" data-event-name="history-sticky-header"><span class="vector-icon mw-ui-icon-wikimedia-history mw-ui-icon-wikimedia-wikimedia-history"></span>

<span></span>
			</a>
			<a href="#" class="cdx-button cdx-button--fake-button cdx-button--fake-button--enabled cdx-button--weight-quiet cdx-button--icon-only mw-watchlink" id="ca-watchstar-sticky-header" tabindex="-1" data-event-name="watch-sticky-header"><span class="vector-icon mw-ui-icon-wikimedia-star mw-ui-icon-wikimedia-wikimedia-star"></span>

<span></span>
			</a>
			<a href="#" class="cdx-button cdx-button--fake-button cdx-button--fake-button--enabled cdx-button--weight-quiet cdx-button--icon-only" id="ca-edit-sticky-header" tabindex="-1" data-event-name="wikitext-edit-sticky-header"><span class="vector-icon mw-ui-icon-wikimedia-wikiText mw-ui-icon-wikimedia-wikimedia-wikiText"></span>

<span></span>
			</a>
			<a href="#" class="cdx-button cdx-button--fake-button cdx-button--fake-button--enabled cdx-button--weight-quiet cdx-button--icon-only" id="ca-ve-edit-sticky-header" tabindex="-1" data-event-name="ve-edit-sticky-header"><span class="vector-icon mw-ui-icon-wikimedia-edit mw-ui-icon-wikimedia-wikimedia-edit"></span>

<span></span>
			</a>
			<a href="#" class="cdx-button cdx-button--fake-button cdx-button--fake-button--enabled cdx-button--weight-quiet cdx-button--icon-only" id="ca-viewsource-sticky-header" tabindex="-1" data-event-name="ve-edit-protected-sticky-header"><span class="vector-icon mw-ui-icon-wikimedia-editLock mw-ui-icon-wikimedia-wikimedia-editLock"></span>

<span></span>
			</a>
		</div>
			<div class="vector-sticky-header-buttons">
				<button class="cdx-button cdx-button--weight-quiet mw-interlanguage-selector" id="p-lang-btn-sticky-header" tabindex="-1" data-event-name="ui.dropdown-p-lang-btn-sticky-header"><span class="vector-icon mw-ui-icon-wikimedia-language mw-ui-icon-wikimedia-wikimedia-language"></span>

<span>30 languages</span>
			</button>
			<a href="#" class="cdx-button cdx-button--fake-button cdx-button--fake-button--enabled cdx-button--weight-quiet cdx-button--action-progressive" id="ca-addsection-sticky-header" tabindex="-1" data-event-name="addsection-sticky-header"><span class="vector-icon mw-ui-icon-speechBubbleAdd-progressive mw-ui-icon-wikimedia-speechBubbleAdd-progressive"></span>

<span>Add topic</span>
			</a>
		</div>
			<div class="vector-sticky-header-icon-end">
				<div class="vector-user-links">
				</div>
			</div>
		</div>
	</div>
</div>
<div class="mw-portlet mw-portlet-dock-bottom emptyPortlet" id="p-dock-bottom">
	<ul>
		
	</ul>
</div>
<script>(RLQ=window.RLQ||[]).push(function(){mw.config.set({"wgHostname":"mw-web.eqiad.main-6fb68b447f-6nrc4","wgBackendResponseTime":148,"wgPageParseReport":{"limitreport":{"cputime":"1.237","walltime":"1.487","ppvisitednodes":{"value":6554,"limit":1000000},"revisionsize":{"value":52124,"limit":2097152},"postexpandincludesize":{"value":389179,"limit":2097152},"templateargumentsize":{"value":84459,"limit":2097152},"expansiondepth":{"value":17,"limit":100},"expensivefunctioncount":{"value":8,"limit":500},"unstrip-depth":{"value":1,"limit":20},"unstrip-size":{"value":289041,"limit":5000000},"entityaccesscount":{"value":1,"limit":500},"timingprofile":["100.00% 1185.956      1 -total"," 44.95%  533.058      1 Template:Reflist"," 35.50%  421.033     54 Template:Cite_web"," 13.71%  162.617     11 Template:Navbox","  9.99%  118.445      1 Template:Infobox_university","  9.77%  115.869      2 Template:Infobox","  8.74%  103.647      1 Template:City,_University_of_London","  5.98%   70.896      1 Template:Navboxes","  5.67%   67.241      1 Template:Short_description","  4.77%   56.576      9 Template:Cite_news"]},"scribunto":{"limitreport-timeusage":{"value":"0.740","limit":"10.000"},"limitreport-memusage":{"value":8078683,"limit":52428800}},"cachereport":{"origin":"mw-web.codfw.canary-86b8f7b789-tcc2m","timestamp":"20250523235530","ttl":2592000,"transientcontent":false}}});});</script>
<script type="application/ld+json">{"@context":"https:\/\/schema.org","@type":"Article","name":"City, University of London","url":"https:\/\/en.wikipedia.org\/wiki\/City,_University_of_London","sameAs":"http:\/\/www.wikidata.org\/entity\/Q1094046","mainEntity":"http:\/\/www.wikidata.org\/entity\/Q1094046","author":{"@type":"Organization","name":"Contributors to Wikimedia projects"},"publisher":{"@type":"Organization","name":"Wikimedia Foundation, Inc.","logo":{"@type":"ImageObject","url":"https:\/\/www.wikimedia.org\/static\/images\/wmf-hor-googpub.png"}},"datePublished":"2004-03-17T22:40:01Z","dateModified":"2025-04-03T05:31:34Z","image":"https:\/\/upload.wikimedia.org\/wikipedia\/commons\/0\/01\/Arms_of_City%2C_University_of_London.svg","headline":"university"}</script>
</body>
</html>
`


