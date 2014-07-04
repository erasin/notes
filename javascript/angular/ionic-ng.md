# ionic javascript


## tabs

用于底部使用

### ion-tabs

Delegate: **$ionicTabsDelegate**

Powers a multi-tabbed interface with a Tab Bar and a set of "pages" that can be tabbed through.

Assign any tabs class or animation class to the element to define its look and feel.

See the ionTab directive's documentation for more details on individual tabs.

Note: do not place ion-tabs inside of an ion-content element; it has been known to cause a certain CSS bug.
Usage

```html
<ion-tabs class="tabs-positive tabs-icon-only">

  <ion-tab title="Home" icon-on="ion-ios7-filing" icon-off="ion-ios7-filing-outline">
    <!-- Tab 1 content -->
  </ion-tab>

  <ion-tab title="About" icon-on="ion-ios7-clock" icon-off="ion-ios7-clock-outline">
    <!-- Tab 2 content -->
  </ion-tab>

  <ion-tab title="Settings" icon-on="ion-ios7-gear" icon-off="ion-ios7-gear-outline">
    <!-- Tab 3 content -->
  </ion-tab>

</ion-tabs>
```


*[Delegate]: 指令