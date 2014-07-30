# Visual Studio 2010 Setup for node-gyp

On Windows XP/Vista/7, [node-gyp requires Python 2.7 and Visual Studio 2010][link1]

According to the readme file in [Microsoft Visual C++ 2010 Service Pack 1 Compiler Update for the Windows SDK 7.1][link2], to ensure that your system has a supported configuration, uninstall the following products and then reinstall them in the order listed:

1. [Visual C++ 2010][visual2010] Express or Visual Studio 2010
2. [Windows SDK 7.1][winsdk71] Note: If you get error on installation, maybe this link will [help][help] you.
3. [Visual Studio 2010 SP1][vs2010]
4. [Visual C++ 2010 SP1 Compiler Update for the Windows SDK 7.1][vcsdk71]


On x64 environments, the last update in the list fixes errors about missing compilers and `error MSB4019: The imported project "C:\Microsoft.Cpp.Default.props" was not found`.

If you experience the error LNK1181 file kernel32.lib not found, try compiling using the Windows SDK 7.1 Command Prompt start menu shortcut.



[link1]: https://github.com/TooTallNate/node-gyp#installation
[link2]: http://www.microsoft.com/en-us/download/details.aspx?id=4422
[visual2010]: http://www.microsoft.com/visualstudio/eng/downloads#d-2010-express
[winsdk71]: http://www.microsoft.com/en-us/download/details.aspx?id=8279
[help]: http://stackoverflow.com/questions/1901279/windows-7-sdk-installation-failure
[vs2010]: http://www.microsoft.com/en-us/download/details.aspx?id=23691
[vcsdk71]: http://www.microsoft.com/en-us/download/details.aspx?id=4422