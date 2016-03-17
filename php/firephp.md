# firephp


## install

使用pear安装

    pear channel-discover pear.firephp.org
    pear install firephp/FirePHPCore

## USE

[source](http://www.firephp.org/HQ/Use.htm)

NOTE: The information below assumes you are using the FirePHPCore library. Please refer to the server libraries section in the wiki for information on how to log messages with 3rd party frameworks and libraries. If you are using FirePHP 1.0 see the FirePHP 1.0 API Reference.

WARNING: Using FirePHP on production sites can expose sensitive information. You must protect the security of your application by disabling FirePHP logging on your live site. You can do this by removing the logging statements before you upload your code or by restricting FirePHP logging to authorized users only.

The reference information below is targeted at PHP5 users. If you use PHP4 please review the additional PHP4 Notes below.

## Object Oriented API

Use the object oriented API when you want to keep the logging statements in your application. It will provide more flexibility and future-proof your code.

Logging is enabled by default. You can disable it with setEnabled(false). Use this method to disable logging on your live site for everyone except authorized users.

There is also an overloaded method equivalent to fb() (see Procedural API) called FB::send().

```php
require_once('FirePHPCore/FirePHP.class.php');
$firephp = FirePHP::getInstance(true);
$firephp-> *
 
require_once('FirePHPCore/fb.php');
FB:: *
 
$firephp->setEnabled(false);  // or FB::
 
FB::send(/* See fb() */);
```

## Procedural API
  
Zc vbnm,./ASDFGHJKL;'
    QWERTYUIOP[]\`1234567890-The procedural API consists of one function. It is fast to type and the function arguments are overloaded.

Use the fb() function for ad-hock debugging and development when you intend to remove the logging statements again.

```php
require_once('FirePHPCore/fb.php');
 
fb($var);
fb($var, 'Label');
fb($var, FirePHP::*);
fb($var, 'Label', FirePHP::*);
```

## Options & Object Filters


maxObjectDepth  Maximum depth to traverse objects.
maxArrayDepth   Maximum depth to traverse arrays.
maxDepth    Maximum depth to traverse mixed arrays/objects.
useNativeJsonEncode Set to FALSE to use JSON encoder included with FirePHPCore instead of json_encode().
includeLineNumbers  Include File and Line information in message.
To exclude specific members when logging objects use setObjectFilter().

See TIP: FirePHP data volume filtering for an in-depth tutorial.

```php
// Defaults:
$options = array('maxObjectDepth' => 5,
                 'maxArrayDepth' => 5,
                 'maxDepth' => 10,
                 'useNativeJsonEncode' => true,
                 'includeLineNumbers' => true);
 
$firephp->getOptions();
$firephp->setOptions($options);
FB::setOptions($options);
 
$firephp->setObjectFilter('ClassName', array('MemberName'));
```

Error, Exception & Assertion Handling

Convert E_WARNING, E_NOTICE, E_USER_ERROR, E_USER_WARNING, E_USER_NOTICE and E_RECOVERABLE_ERROR errors to ErrorExceptions and send all Exceptions to Firebug automatically if desired.

Assertion errors can be converted to exceptions and thrown if desired.

You can also manually send caught exceptions to Firebug.

```php
$firephp->registerErrorHandler(
            $throwErrorExceptions=false);
$firephp->registerExceptionHandler();
$firephp->registerAssertionHandler(
            $convertAssertionErrorsToExceptions=true,
            $throwAssertionExceptions=false);
 
try {
  throw new Exception('Test Exception');
} catch(Exception $e) {
  $firephp->error($e);  // or FB::
}
```

## Groups

In many cases it is useful to group logging messages together. Groups can be nested programatically and expanded/contracted by the user.

By default groups are expended in the firebug console.

You can change the color of the group label by specyfying a standard HTML color value.

```php
$firephp->group('Test Group');
$firephp->log('Hello World');
$firephp->groupEnd();
 
$firephp->group('Collapsed and Colored Group',
                array('Collapsed' => true,
                      'Color' => '#FF00FF'));
```

## Logging Messages

Priority Styling
These logging methods follow the four Firebug logging priorities. Add an optional label as a second argument to any of these methods.

If you are using the fb() function use the FirePHP::LOG, FirePHP::INFO, FirePHP::WARN, FirePHP::ERROR constants.

```php
$firephp->log('Plain Message');     // or FB::
$firephp->info('Info Message');     // or FB::
$firephp->warn('Warn Message');     // or FB::
$firephp->error('Error Message');   // or FB::
 
$firephp->log('Message','Optional Label');
 
$firephp->fb('Message', FirePHP::*);
```

## Tables
You can log tables of information. Firebug will display the Table Label and allow the user to toggle the display of the table. The first row of the table is automatically used as the heading and the number of columns is dynamically determined.

If you are using the fb() function use the FirePHP::TABLE constant.

```php
$table   = array();
$table[] = array('Col 1 Heading','Col 2 Heading');
$table[] = array('Row 1 Col 1','Row 1 Col 2');
$table[] = array('Row 2 Col 1','Row 2 Col 2');
$table[] = array('Row 3 Col 1','Row 3 Col 2');
 
$firephp->table('Table Label', $table);  // or FB::
 
fb($table, 'Table Label', FirePHP::TABLE);
```

## Traces
You can send a backtrace showing File, Line, Class, Method and Function information including Arguments to clearly show the execution path up to the point in your code where you triggered the trace.

If you are using the fb() function use the FirePHP::TRACE constant.

```php
$firephp->trace('Trace Label');  // or FB::
 
fb('Trace Label', FirePHP::TRACE);
```

## PHP4 Notes

Make sure to use the PHP4 versions of the library files.

FirePHP::getInstance() auto creates a global FirePHP object called $FirePHP_Instance and returns a reference to it. Make sure to remember the & symbol when assigning the returned reference to ensure the object is not copied.

All FirePHP::* class constants are accessible via global FirePHP_* constants.

Exception handling functionality will not work as there is no exception support in PHP4. registerErrorHandler() always logs errors to FirePHP.

```php
require_once('FirePHPCore/FirePHP.class.php4');
require_once('FirePHPCore/fb.php4');
 
$firephp =& FirePHP::getInstance(true);
 
fb($var, FirePHP_*);
```
