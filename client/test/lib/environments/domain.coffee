helpers = require '../helpers/helpers.js'
assert  = require 'assert'
environmentHelpers = require '../helpers/environmenthelpers.js'

domainItem  = '.kdlistitemview-domain:last-child'
loader      = domainItem + '.in-progress'


module.exports =


  # openVmDomain: (browser) ->

  #   ## This test is disabled because DNS propagation takes time and we are running test before DNS propagation complete.

  #   modalSelector = '.activity-modal.vm-settings'
  #   linkSelector  = modalSelector + ' .assigned-url .custom-link-view'

  #   helpers.beginTest(browser)
  #   helpers.waitForVMRunning(browser)

  #   environmentHelpers.openVmSettingsModal(browser)

  #   browser
  #     .waitForElementVisible   linkSelector, 20000
  #     .getAttribute            linkSelector, 'href', (result) ->
  #       domain = result.value

  #       browser
  #         .url domain
  #         .pause  5000
  #         .waitForElementVisible  '#container', 20000
  #         .waitForElementVisible  '#container .hellobox', 20000
  #         .end()


  addDomain: (browser) ->

    environmentHelpers.addDomain(browser)
    browser.end()


  deleteDomain: (browser) ->

    domainName = environmentHelpers.addDomain(browser)

    browser
      .moveToElement             domainItem, 10, 10
      .click                     domainItem + ' span.remove'
      .waitForElementVisible     loader, 10000
      .waitForElementNotVisible  loader, 20000
      .getText                   domainItem, (result) =>
        assert.notEqual          result.value, domainName # Assertion

        browser.end()


  assignDomain: (browser) ->

    domainName = environmentHelpers.addDomain(browser)

    browser
      .waitForElementVisible     domainItem, 20000
      .click                     domainItem + ' .koding-on-off.on'
      .waitForElementVisible     '.in-progress', 10000
      .waitForElementNotPresent  '.in-progress', 20000
      .refresh()

    environmentHelpers.openDomainSettings(browser)

    browser
      .waitForElementVisible     domainItem + ' .koding-on-off.off', 20000
      .end()


  addInvalidDomain: (browser) ->

    helpers.beginTest(browser)
    helpers.waitForVMRunning(browser)

    invalidDomainName    = ';invalidVmName;'
    buttonSelector       = '.domains .kdheaderview button.add-button'
    buttonLoaderSelector = '.add-view button.loading'
    errorSelector        = '.kdmodal-inner .domains .notification.error'

    environmentHelpers.openDomainSettings(browser)

    browser
    .waitForElementVisible    buttonSelector, 20000
    .click                    buttonSelector
    .waitForElementVisible    '.add-view input.hitenterview', 20000
    .setValue                 '.add-view input.hitenterview', invalidDomainName + '\n'
    .waitForElementVisible    buttonLoaderSelector, 10000
    .waitForElementNotPresent buttonLoaderSelector, 20000
    .waitForElementVisible    errorSelector, 20000 # Assertion
    .end()


