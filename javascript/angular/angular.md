# angular note

<https://thinkster.io/angulartutorial/a-better-way-to-learn-angularjs/>

<http://www.ng-newsletter.com/posts/angular-for-the-jquery-developer.html>

animate

指令 | 支持动画 
-----|------------------------------
Directive |  Supported Animations
ngRepeat |   enter, leave and move
ngView | enter and leave
ngInclude |  enter and leave
ngSwitch |   enter and leave
ngIf |   enter and leave
ngClass | add and remove (the CSS class(es) present)
ngShow & ngHide | add and remove (the ng-hide class value)
form & ngModel | add and remove (dirty, pristine, valid, invalid & all other validations)
ngMessages | add and remove (ng-active & ng-inactive)
ngMessage  | enter and leave


## 坑

form 表单，表单内验证的name属性为form的属性，并非为`$ng-model`绑定的变量。

提交方法使用参数传值。直接使用`$scope`来传值有可能出错，多为`$promise`所产生的问题。

数据整理多使用服务器端。js客户端做少量处理，多做组合。
